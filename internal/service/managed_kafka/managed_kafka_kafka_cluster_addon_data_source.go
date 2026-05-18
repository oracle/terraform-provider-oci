// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package managed_kafka

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_managed_kafka "github.com/oracle/oci-go-sdk/v65/managedkafka"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ManagedKafkaKafkaClusterAddonDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["kafka_cluster_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(ManagedKafkaKafkaClusterAddonResource(), fieldMap, readSingularManagedKafkaKafkaClusterAddonWithContext)
}

func readSingularManagedKafkaKafkaClusterAddonWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ManagedKafkaKafkaClusterAddonDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KafkaClusterClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type ManagedKafkaKafkaClusterAddonDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_managed_kafka.KafkaClusterClient
	Res    *oci_managed_kafka.GetAddonResponse
}

func (s *ManagedKafkaKafkaClusterAddonDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ManagedKafkaKafkaClusterAddonDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_managed_kafka.GetAddonRequest{}

	if addonName, ok := s.D.GetOkExists("name"); ok {
		tmp := addonName.(string)
		request.AddonName = &tmp
	}

	if kafkaClusterId, ok := s.D.GetOkExists("kafka_cluster_id"); ok {
		tmp := kafkaClusterId.(string)
		request.KafkaClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "managed_kafka")

	response, err := s.Client.GetAddon(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ManagedKafkaKafkaClusterAddonDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ManagedKafkaKafkaClusterAddonDataSource-", ManagedKafkaKafkaClusterAddonDataSource(), s.D))
	switch v := (s.Res.KafkaClusterAddon).(type) {
	case oci_managed_kafka.PublicConnectivityAddon:
		s.D.Set("addon_type", "PUBLICCONNECTIVITY")

		s.D.Set("authentication_mechanism", v.AuthenticationMechanism)

		if v.BootstrapUrl != nil {
			s.D.Set("bootstrap_url", *v.BootstrapUrl)
		}

		s.D.Set("network_cidrs", v.NetworkCidrs)

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.Name != nil {
			s.D.Set("name", *v.Name)
		}

		s.D.Set("state", v.LifecycleState)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'addon_type' of unknown type %v", s.Res.KafkaClusterAddon)
		return nil
	}

	return nil
}
