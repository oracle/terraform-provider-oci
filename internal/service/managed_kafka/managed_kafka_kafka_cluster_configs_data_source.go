// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package managed_kafka

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_managed_kafka "github.com/oracle/oci-go-sdk/v65/managedkafka"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ManagedKafkaKafkaClusterConfigsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readManagedKafkaKafkaClusterConfigs,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"kafka_cluster_config_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ManagedKafkaKafkaClusterConfigResource()),
						},
					},
				},
			},
		},
	}
}

func readManagedKafkaKafkaClusterConfigs(d *schema.ResourceData, m interface{}) error {
	sync := &ManagedKafkaKafkaClusterConfigsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KafkaClusterClient()

	return tfresource.ReadResource(sync)
}

type ManagedKafkaKafkaClusterConfigsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_managed_kafka.KafkaClusterClient
	Res    *oci_managed_kafka.ListKafkaClusterConfigsResponse
}

func (s *ManagedKafkaKafkaClusterConfigsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagedKafkaKafkaClusterConfigsDataSourceCrud) Get() error {
	request := oci_managed_kafka.ListKafkaClusterConfigsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_managed_kafka.KafkaClusterConfigLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "managed_kafka")

	response, err := s.Client.ListKafkaClusterConfigs(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListKafkaClusterConfigs(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ManagedKafkaKafkaClusterConfigsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ManagedKafkaKafkaClusterConfigsDataSource-", ManagedKafkaKafkaClusterConfigsDataSource(), s.D))
	resources := []map[string]interface{}{}
	kafkaClusterConfig := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, KafkaClusterConfigSummaryToMap(item))
	}
	kafkaClusterConfig["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ManagedKafkaKafkaClusterConfigsDataSource().Schema["kafka_cluster_config_collection"].Elem.(*schema.Resource).Schema)
		kafkaClusterConfig["items"] = items
	}

	resources = append(resources, kafkaClusterConfig)
	if err := s.D.Set("kafka_cluster_config_collection", resources); err != nil {
		return err
	}

	return nil
}
