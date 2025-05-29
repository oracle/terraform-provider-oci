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

func ManagedKafkaKafkaClusterDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["kafka_cluster_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ManagedKafkaKafkaClusterResource(), fieldMap, readSingularManagedKafkaKafkaCluster)
}

func readSingularManagedKafkaKafkaCluster(d *schema.ResourceData, m interface{}) error {
	sync := &ManagedKafkaKafkaClusterDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KafkaClusterClient()

	return tfresource.ReadResource(sync)
}

type ManagedKafkaKafkaClusterDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_managed_kafka.KafkaClusterClient
	Res    *oci_managed_kafka.GetKafkaClusterResponse
}

func (s *ManagedKafkaKafkaClusterDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagedKafkaKafkaClusterDataSourceCrud) Get() error {
	request := oci_managed_kafka.GetKafkaClusterRequest{}

	if kafkaClusterId, ok := s.D.GetOkExists("kafka_cluster_id"); ok {
		tmp := kafkaClusterId.(string)
		request.KafkaClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "managed_kafka")

	response, err := s.Client.GetKafkaCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ManagedKafkaKafkaClusterDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	accessSubnets := []interface{}{}
	for _, item := range s.Res.AccessSubnets {
		accessSubnets = append(accessSubnets, SubnetSetToMap(item))
	}
	s.D.Set("access_subnets", accessSubnets)

	if s.Res.BrokerShape != nil {
		s.D.Set("broker_shape", []interface{}{BrokerShapeToMap(s.Res.BrokerShape)})
	} else {
		s.D.Set("broker_shape", nil)
	}

	if s.Res.ClientCertificateBundle != nil {
		s.D.Set("client_certificate_bundle", *s.Res.ClientCertificateBundle)
	}

	if s.Res.ClusterConfigId != nil {
		s.D.Set("cluster_config_id", *s.Res.ClusterConfigId)
	}

	if s.Res.ClusterConfigVersion != nil {
		s.D.Set("cluster_config_version", *s.Res.ClusterConfigVersion)
	}

	s.D.Set("cluster_type", s.Res.ClusterType)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("coordination_type", s.Res.CoordinationType)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	kafkaBootstrapUrls := []interface{}{}
	for _, item := range s.Res.KafkaBootstrapUrls {
		kafkaBootstrapUrls = append(kafkaBootstrapUrls, BootstrapUrlToMap(item))
	}
	s.D.Set("kafka_bootstrap_urls", kafkaBootstrapUrls)

	if s.Res.KafkaVersion != nil {
		s.D.Set("kafka_version", *s.Res.KafkaVersion)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.SecretId != nil {
		s.D.Set("secret_id", *s.Res.SecretId)
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
