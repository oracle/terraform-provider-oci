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

func QueueQueueDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["queue_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(QueueQueueResource(), fieldMap, readSingularQueueQueue)
}

func readSingularQueueQueue(d *schema.ResourceData, m interface{}) error {
	sync := &QueueQueueDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).QueueAdminClient()

	return tfresource.ReadResource(sync)
}

type QueueQueueDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_queue.QueueAdminClient
	Res    *oci_queue.GetQueueResponse
}

func (s *QueueQueueDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *QueueQueueDataSourceCrud) Get() error {
	request := oci_queue.GetQueueRequest{}

	if queueId, ok := s.D.GetOkExists("queue_id"); ok {
		tmp := queueId.(string)
		request.QueueId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "queue")

	response, err := s.Client.GetQueue(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *QueueQueueDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ChannelConsumptionLimit != nil {
		s.D.Set("channel_consumption_limit", *s.Res.ChannelConsumptionLimit)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CustomEncryptionKeyId != nil {
		s.D.Set("custom_encryption_key_id", *s.Res.CustomEncryptionKeyId)
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

	if s.Res.MessagesEndpoint != nil {
		s.D.Set("messages_endpoint", *s.Res.MessagesEndpoint)
	}

	if s.Res.RetentionInSeconds != nil {
		s.D.Set("retention_in_seconds", *s.Res.RetentionInSeconds)
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

	if s.Res.TimeoutInSeconds != nil {
		s.D.Set("timeout_in_seconds", *s.Res.TimeoutInSeconds)
	}

	if s.Res.VisibilityInSeconds != nil {
		s.D.Set("visibility_in_seconds", *s.Res.VisibilityInSeconds)
	}

	return nil
}
