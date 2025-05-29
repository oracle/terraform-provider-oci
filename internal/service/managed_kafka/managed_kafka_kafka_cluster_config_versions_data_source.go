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

func ManagedKafkaKafkaClusterConfigVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readManagedKafkaKafkaClusterConfigVersions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"kafka_cluster_config_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"kafka_cluster_config_version_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"config_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"version_number": {
										Type:     schema.TypeInt,
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

func readManagedKafkaKafkaClusterConfigVersions(d *schema.ResourceData, m interface{}) error {
	sync := &ManagedKafkaKafkaClusterConfigVersionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KafkaClusterClient()

	return tfresource.ReadResource(sync)
}

type ManagedKafkaKafkaClusterConfigVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_managed_kafka.KafkaClusterClient
	Res    *oci_managed_kafka.ListKafkaClusterConfigVersionsResponse
}

func (s *ManagedKafkaKafkaClusterConfigVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagedKafkaKafkaClusterConfigVersionsDataSourceCrud) Get() error {
	request := oci_managed_kafka.ListKafkaClusterConfigVersionsRequest{}

	if kafkaClusterConfigId, ok := s.D.GetOkExists("kafka_cluster_config_id"); ok {
		tmp := kafkaClusterConfigId.(string)
		request.KafkaClusterConfigId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "managed_kafka")

	response, err := s.Client.ListKafkaClusterConfigVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListKafkaClusterConfigVersions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ManagedKafkaKafkaClusterConfigVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ManagedKafkaKafkaClusterConfigVersionsDataSource-", ManagedKafkaKafkaClusterConfigVersionsDataSource(), s.D))
	resources := []map[string]interface{}{}
	kafkaClusterConfigVersion := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, KafkaClusterConfigVersionSummaryToMap(item))
	}
	kafkaClusterConfigVersion["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ManagedKafkaKafkaClusterConfigVersionsDataSource().Schema["kafka_cluster_config_version_collection"].Elem.(*schema.Resource).Schema)
		kafkaClusterConfigVersion["items"] = items
	}

	resources = append(resources, kafkaClusterConfigVersion)
	if err := s.D.Set("kafka_cluster_config_version_collection", resources); err != nil {
		return err
	}

	return nil
}

func KafkaClusterConfigVersionSummaryToMap(obj oci_managed_kafka.KafkaClusterConfigVersionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConfigId != nil {
		result["config_id"] = string(*obj.ConfigId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.VersionNumber != nil {
		result["version_number"] = int(*obj.VersionNumber)
	}

	return result
}
