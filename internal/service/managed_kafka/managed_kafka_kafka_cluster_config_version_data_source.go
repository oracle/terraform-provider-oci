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

func ManagedKafkaKafkaClusterConfigVersionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularManagedKafkaKafkaClusterConfigVersion,
		Schema: map[string]*schema.Schema{
			"kafka_cluster_config_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"version_number": {
				Type:     schema.TypeInt,
				Required: true,
			},
			// Computed
			"config_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"properties": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularManagedKafkaKafkaClusterConfigVersion(d *schema.ResourceData, m interface{}) error {
	sync := &ManagedKafkaKafkaClusterConfigVersionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KafkaClusterClient()

	return tfresource.ReadResource(sync)
}

type ManagedKafkaKafkaClusterConfigVersionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_managed_kafka.KafkaClusterClient
	Res    *oci_managed_kafka.GetKafkaClusterConfigVersionResponse
}

func (s *ManagedKafkaKafkaClusterConfigVersionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagedKafkaKafkaClusterConfigVersionDataSourceCrud) Get() error {
	request := oci_managed_kafka.GetKafkaClusterConfigVersionRequest{}

	if kafkaClusterConfigId, ok := s.D.GetOkExists("kafka_cluster_config_id"); ok {
		tmp := kafkaClusterConfigId.(string)
		request.KafkaClusterConfigId = &tmp
	}

	if versionNumber, ok := s.D.GetOkExists("version_number"); ok {
		tmp := versionNumber.(int)
		request.VersionNumber = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "managed_kafka")

	response, err := s.Client.GetKafkaClusterConfigVersion(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ManagedKafkaKafkaClusterConfigVersionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ManagedKafkaKafkaClusterConfigVersionDataSource-", ManagedKafkaKafkaClusterConfigVersionDataSource(), s.D))

	if s.Res.ConfigId != nil {
		s.D.Set("config_id", *s.Res.ConfigId)
	}

	s.D.Set("properties", s.Res.Properties)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
