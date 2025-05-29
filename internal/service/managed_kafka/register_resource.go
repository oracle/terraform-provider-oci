// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package managed_kafka

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterResource() {
	tfresource.RegisterResource("oci_managed_kafka_kafka_cluster", ManagedKafkaKafkaClusterResource())
	tfresource.RegisterResource("oci_managed_kafka_kafka_cluster_config", ManagedKafkaKafkaClusterConfigResource())
	tfresource.RegisterResource("oci_managed_kafka_kafka_cluster_superusers_management", ManagedKafkaKafkaClusterSuperusersManagementResource())
}
