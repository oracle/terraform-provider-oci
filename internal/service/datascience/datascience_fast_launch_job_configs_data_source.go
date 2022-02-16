// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v58/datascience"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func DatascienceFastLaunchJobConfigsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatascienceFastLaunchJobConfigs,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fast_launch_job_configs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"core_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"managed_egress_support": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"memory_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shape_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"shape_series": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatascienceFastLaunchJobConfigs(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceFastLaunchJobConfigsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceFastLaunchJobConfigsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.ListFastLaunchJobConfigsResponse
}

func (s *DatascienceFastLaunchJobConfigsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceFastLaunchJobConfigsDataSourceCrud) Get() error {
	request := oci_datascience.ListFastLaunchJobConfigsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.ListFastLaunchJobConfigs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListFastLaunchJobConfigs(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatascienceFastLaunchJobConfigsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatascienceFastLaunchJobConfigsDataSource-", DatascienceFastLaunchJobConfigsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		fastLaunchJobConfig := map[string]interface{}{}

		if r.CoreCount != nil {
			fastLaunchJobConfig["core_count"] = *r.CoreCount
		}

		fastLaunchJobConfig["managed_egress_support"] = r.ManagedEgressSupport

		if r.MemoryInGBs != nil {
			fastLaunchJobConfig["memory_in_gbs"] = *r.MemoryInGBs
		}

		if r.Name != nil {
			fastLaunchJobConfig["name"] = *r.Name
		}

		if r.ShapeName != nil {
			fastLaunchJobConfig["shape_name"] = *r.ShapeName
		}

		fastLaunchJobConfig["shape_series"] = r.ShapeSeries

		resources = append(resources, fastLaunchJobConfig)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatascienceFastLaunchJobConfigsDataSource().Schema["fast_launch_job_configs"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("fast_launch_job_configs", resources); err != nil {
		return err
	}

	return nil
}
