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

func QueueConsumerGroupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["consumer_group_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(QueueConsumerGroupResource(), fieldMap, readSingularQueueConsumerGroup)
}

func readSingularQueueConsumerGroup(d *schema.ResourceData, m interface{}) error {
	sync := &QueueConsumerGroupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).QueueAdminClient()

	return tfresource.ReadResource(sync)
}

type QueueConsumerGroupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_queue.QueueAdminClient
	Res    *oci_queue.GetConsumerGroupResponse
}

func (s *QueueConsumerGroupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *QueueConsumerGroupDataSourceCrud) Get() error {
	request := oci_queue.GetConsumerGroupRequest{}

	if consumerGroupId, ok := s.D.GetOkExists("consumer_group_id"); ok {
		tmp := consumerGroupId.(string)
		request.ConsumerGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "queue")

	response, err := s.Client.GetConsumerGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *QueueConsumerGroupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Filter != nil {
		s.D.Set("consumer_group_filter", *s.Res.Filter)
	}

	if s.Res.DeadLetterQueueDeliveryCount != nil {
		s.D.Set("dead_letter_queue_delivery_count", *s.Res.DeadLetterQueueDeliveryCount)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.QueueId != nil {
		s.D.Set("queue_id", *s.Res.QueueId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
