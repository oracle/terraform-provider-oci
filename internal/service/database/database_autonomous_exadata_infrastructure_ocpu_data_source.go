// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v58/database"
)

func DatabaseAutonomousExadataInfrastructureOcpuDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseAutonomousExadataInfrastructureOcpu,
		Schema: map[string]*schema.Schema{
			"autonomous_exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"by_workload_type": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"adw": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"atp": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
			"consumed_cpu": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"total_cpu": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseAutonomousExadataInfrastructureOcpu(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousExadataInfrastructureOcpuDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousExadataInfrastructureOcpuDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetExadataInfrastructureOcpusResponse
}

func (s *DatabaseAutonomousExadataInfrastructureOcpuDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousExadataInfrastructureOcpuDataSourceCrud) Get() error {
	request := oci_database.GetExadataInfrastructureOcpusRequest{}

	if autonomousExadataInfrastructureId, ok := s.D.GetOkExists("autonomous_exadata_infrastructure_id"); ok {
		tmp := autonomousExadataInfrastructureId.(string)
		request.AutonomousExadataInfrastructureId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetExadataInfrastructureOcpus(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseAutonomousExadataInfrastructureOcpuDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousExadataInfrastructureOcpuDataSource-", DatabaseAutonomousExadataInfrastructureOcpuDataSource(), s.D))

	if s.Res.ByWorkloadType != nil {
		s.D.Set("by_workload_type", []interface{}{WorkloadTypeToMap(s.Res.ByWorkloadType)})
	} else {
		s.D.Set("by_workload_type", nil)
	}

	if s.Res.ConsumedCpu != nil {
		s.D.Set("consumed_cpu", *s.Res.ConsumedCpu)
	}

	if s.Res.TotalCpu != nil {
		s.D.Set("total_cpu", *s.Res.TotalCpu)
	}

	return nil
}

func WorkloadTypeToMap(obj *oci_database.WorkloadType) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Adw != nil {
		result["adw"] = float32(*obj.Adw)
	}

	if obj.Atp != nil {
		result["atp"] = float32(*obj.Atp)
	}

	return result
}
