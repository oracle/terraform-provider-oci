// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package em_warehouse

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_em_warehouse "github.com/oracle/oci-go-sdk/v65/emwarehouse"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func EmWarehouseEmWarehouseResourceUsageDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularEmWarehouseEmWarehouseResourceUsage,
		Schema: map[string]*schema.Schema{
			"em_warehouse_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"em_instance_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"em_instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"em_discoverer_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"em_host": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"em_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"targets_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"operations_insights_warehouse_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"schema_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"targets_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func readSingularEmWarehouseEmWarehouseResourceUsage(d *schema.ResourceData, m interface{}) error {
	sync := &EmWarehouseEmWarehouseResourceUsageDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmWarehouseClient()

	return tfresource.ReadResource(sync)
}

type EmWarehouseEmWarehouseResourceUsageDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_em_warehouse.EmWarehouseClient
	Res    *oci_em_warehouse.GetEmWarehouseResourceUsageResponse
}

func (s *EmWarehouseEmWarehouseResourceUsageDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *EmWarehouseEmWarehouseResourceUsageDataSourceCrud) Get() error {
	request := oci_em_warehouse.GetEmWarehouseResourceUsageRequest{}

	if emWarehouseId, ok := s.D.GetOkExists("em_warehouse_id"); ok {
		tmp := emWarehouseId.(string)
		request.EmWarehouseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "emwarehouse")

	response, err := s.Client.GetEmWarehouseResourceUsage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *EmWarehouseEmWarehouseResourceUsageDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.EmInstanceCount != nil {
		s.D.Set("em_instance_count", *s.Res.EmInstanceCount)
	}

	emInstances := []interface{}{}
	for _, item := range s.Res.EmInstances {
		emInstances = append(emInstances, EmInstancesDetailsToMap(item))
	}
	s.D.Set("em_instances", emInstances)

	if s.Res.OperationsInsightsWarehouseId != nil {
		s.D.Set("operations_insights_warehouse_id", *s.Res.OperationsInsightsWarehouseId)
	}

	if s.Res.SchemaName != nil {
		s.D.Set("schema_name", *s.Res.SchemaName)
	}

	if s.Res.TargetsCount != nil {
		s.D.Set("targets_count", *s.Res.TargetsCount)
	}

	return nil
}

func EmInstancesDetailsToMap(obj oci_em_warehouse.EmInstancesDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EmDiscovererUrl != nil {
		result["em_discoverer_url"] = string(*obj.EmDiscovererUrl)
	}

	if obj.EmHost != nil {
		result["em_host"] = string(*obj.EmHost)
	}

	if obj.EmId != nil {
		result["em_id"] = string(*obj.EmId)
	}

	if obj.TargetsCount != nil {
		result["targets_count"] = int(*obj.TargetsCount)
	}

	return result
}
