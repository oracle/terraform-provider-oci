// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeDiscoveryAnalyticDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDataSafeDiscoveryAnalytic,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"group_by": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sensitive_data_model_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"count": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dimensions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"sensitive_data_model_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"metric_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_data_safe_discovery_analytic", "oci_data_safe_discovery_analytics"),
	}
}

func readSingularDataSafeDiscoveryAnalytic(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeDiscoveryAnalyticDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeDiscoveryAnalyticDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListDiscoveryAnalyticsResponse
}

func (s *DataSafeDiscoveryAnalyticDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeDiscoveryAnalyticDataSourceCrud) Get() error {
	request := oci_data_safe.ListDiscoveryAnalyticsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if groupBy, ok := s.D.GetOkExists("group_by"); ok {
		request.GroupBy = oci_data_safe.ListDiscoveryAnalyticsGroupByEnum(groupBy.(string))
	}

	if sensitiveDataModelId, ok := s.D.GetOkExists("sensitive_data_model_id"); ok {
		tmp := sensitiveDataModelId.(string)
		request.SensitiveDataModelId = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListDiscoveryAnalytics(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataSafeDiscoveryAnalyticDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeDiscoveryAnalyticDataSource-", DataSafeDiscoveryAnalyticDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DiscoveryAnalyticsSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
