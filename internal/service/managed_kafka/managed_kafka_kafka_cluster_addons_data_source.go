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

func ManagedKafkaKafkaClusterAddonsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readManagedKafkaKafkaClusterAddonsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"kafka_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"addon_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     ManagedKafkaKafkaClusterAddonResource(),
						},
					},
				},
			},
		},
	}
}

func readManagedKafkaKafkaClusterAddonsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ManagedKafkaKafkaClusterAddonsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KafkaClusterClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type ManagedKafkaKafkaClusterAddonsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_managed_kafka.KafkaClusterClient
	Res    *oci_managed_kafka.ListAddonsResponse
}

func (s *ManagedKafkaKafkaClusterAddonsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagedKafkaKafkaClusterAddonsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_managed_kafka.ListAddonsRequest{}

	if kafkaClusterId, ok := s.D.GetOkExists("kafka_cluster_id"); ok {
		tmp := kafkaClusterId.(string)
		request.KafkaClusterId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_managed_kafka.KafkaClusterAddonLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "managed_kafka")

	response, err := s.Client.ListAddons(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAddons(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ManagedKafkaKafkaClusterAddonsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ManagedKafkaKafkaClusterAddonsDataSource-", ManagedKafkaKafkaClusterAddonsDataSource(), s.D))
	resources := []map[string]interface{}{}
	kafkaClusterAddon := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AddonSummaryToMap(item))
	}
	kafkaClusterAddon["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ManagedKafkaKafkaClusterAddonsDataSource().Schema["addon_collection"].Elem.(*schema.Resource).Schema)
		kafkaClusterAddon["items"] = items
	}

	resources = append(resources, kafkaClusterAddon)
	if err := s.D.Set("addon_collection", resources); err != nil {
		return err
	}

	return nil
}
