// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package managed_kafka

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_managed_kafka "github.com/oracle/oci-go-sdk/v65/managedkafka"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ManagedKafkaNodeShapesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readManagedKafkaNodeShapesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"node_shape_collection": {
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
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readManagedKafkaNodeShapesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ManagedKafkaNodeShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KafkaClusterClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type ManagedKafkaNodeShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_managed_kafka.KafkaClusterClient
	Res    *oci_managed_kafka.ListNodeShapesResponse
}

func (s *ManagedKafkaNodeShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagedKafkaNodeShapesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_managed_kafka.ListNodeShapesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "managed_kafka")

	response, err := s.Client.ListNodeShapes(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListNodeShapes(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ManagedKafkaNodeShapesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ManagedKafkaNodeShapesDataSource-", ManagedKafkaNodeShapesDataSource(), s.D))
	resources := []map[string]interface{}{}
	nodeShape := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, NodeShapeSummaryToMap(item))
	}
	nodeShape["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ManagedKafkaNodeShapesDataSource().Schema["node_shape_collection"].Elem.(*schema.Resource).Schema)
		nodeShape["items"] = items
	}

	resources = append(resources, nodeShape)
	if err := s.D.Set("node_shape_collection", resources); err != nil {
		return err
	}

	return nil
}

func NodeShapeSummaryToMap(obj oci_managed_kafka.NodeShapeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}
