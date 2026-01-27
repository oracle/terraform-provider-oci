// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package queue

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_queue "github.com/oracle/oci-go-sdk/v65/queue"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func QueueConsumerGroupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readQueueConsumerGroups,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"queue_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"consumer_group_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(QueueConsumerGroupResource()),
						},
					},
				},
			},
		},
	}
}

func readQueueConsumerGroups(d *schema.ResourceData, m interface{}) error {
	sync := &QueueConsumerGroupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).QueueAdminClient()

	return tfresource.ReadResource(sync)
}

type QueueConsumerGroupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_queue.QueueAdminClient
	Res    *oci_queue.ListConsumerGroupsResponse
}

func (s *QueueConsumerGroupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *QueueConsumerGroupsDataSourceCrud) Get() error {
	request := oci_queue.ListConsumerGroupsRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if queueId, ok := s.D.GetOkExists("queue_id"); ok {
		tmp := queueId.(string)
		request.QueueId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_queue.ConsumerGroupLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "queue")

	response, err := s.Client.ListConsumerGroups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListConsumerGroups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *QueueConsumerGroupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("QueueConsumerGroupsDataSource-", QueueConsumerGroupsDataSource(), s.D))
	resources := []map[string]interface{}{}
	consumerGroup := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ConsumerGroupSummaryToMap(item))
	}
	consumerGroup["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, QueueConsumerGroupsDataSource().Schema["consumer_group_collection"].Elem.(*schema.Resource).Schema)
		consumerGroup["items"] = items
	}

	resources = append(resources, consumerGroup)
	if err := s.D.Set("consumer_group_collection", resources); err != nil {
		return err
	}

	return nil
}
