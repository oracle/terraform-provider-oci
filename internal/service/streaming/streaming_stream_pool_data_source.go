// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package streaming

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_streaming "github.com/oracle/oci-go-sdk/v58/streaming"
)

func StreamingStreamPoolDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["stream_pool_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(StreamingStreamPoolResource(), fieldMap, readSingularStreamingStreamPool)
}

func readSingularStreamingStreamPool(d *schema.ResourceData, m interface{}) error {
	sync := &StreamingStreamPoolDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StreamAdminClient()

	return tfresource.ReadResource(sync)
}

type StreamingStreamPoolDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_streaming.StreamAdminClient
	Res    *oci_streaming.GetStreamPoolResponse
}

func (s *StreamingStreamPoolDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StreamingStreamPoolDataSourceCrud) Get() error {
	request := oci_streaming.GetStreamPoolRequest{}

	if streamPoolId, ok := s.D.GetOkExists("stream_pool_id"); ok {
		tmp := streamPoolId.(string)
		request.StreamPoolId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "streaming")

	response, err := s.Client.GetStreamPool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *StreamingStreamPoolDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CustomEncryptionKey != nil {
		s.D.Set("custom_encryption_key", []interface{}{CustomEncryptionKeyToMap(s.Res.CustomEncryptionKey)})
	} else {
		s.D.Set("custom_encryption_key", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.EndpointFqdn != nil {
		s.D.Set("endpoint_fqdn", *s.Res.EndpointFqdn)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsPrivate != nil {
		s.D.Set("is_private", *s.Res.IsPrivate)
	}

	if s.Res.KafkaSettings != nil {
		s.D.Set("kafka_settings", []interface{}{KafkaSettingsToMap(s.Res.KafkaSettings)})
	} else {
		s.D.Set("kafka_settings", nil)
	}

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.PrivateEndpointSettings != nil {
		s.D.Set("private_endpoint_settings", []interface{}{PrivateEndpointSettingsToMap(s.Res.PrivateEndpointSettings, true)})
	} else {
		s.D.Set("private_endpoint_settings", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
