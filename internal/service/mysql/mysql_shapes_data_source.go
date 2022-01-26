// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package mysql

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_mysql "github.com/oracle/oci-go-sdk/v56/mysql"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func MysqlShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMysqlShapes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_supported_for": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"shapes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"cpu_core_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"is_supported_for": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"memory_size_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readMysqlShapes(d *schema.ResourceData, m interface{}) error {
	sync := &MysqlShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MysqlaasClient()

	return tfresource.ReadResource(sync)
}

type MysqlShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_mysql.MysqlaasClient
	Res    *oci_mysql.ListShapesResponse
}

func (s *MysqlShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MysqlShapesDataSourceCrud) Get() error {
	request := oci_mysql.ListShapesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if isSupportedFor, ok := s.D.GetOkExists("is_supported_for"); ok {
		interfaces := isSupportedFor.([]interface{})
		tmp := make([]oci_mysql.ListShapesIsSupportedForEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_mysql.ListShapesIsSupportedForEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("is_supported_for") {
			request.IsSupportedFor = tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "mysql")

	response, err := s.Client.ListShapes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MysqlShapesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MysqlShapesDataSource-", MysqlShapesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		shape := map[string]interface{}{}

		if r.CpuCoreCount != nil {
			shape["cpu_core_count"] = *r.CpuCoreCount
		}

		shape["is_supported_for"] = r.IsSupportedFor

		if r.MemorySizeInGBs != nil {
			shape["memory_size_in_gbs"] = *r.MemorySizeInGBs
		}

		if r.Name != nil {
			shape["name"] = *r.Name
		}

		resources = append(resources, shape)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, MysqlShapesDataSource().Schema["shapes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("shapes", resources); err != nil {
		return err
	}

	return nil
}
