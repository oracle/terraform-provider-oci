// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package managed_kafka

import "github.com/oracle/terraform-provider-oci/internal/tfresource"

func RegisterDatasource() {
	tfresource.RegisterDatasource("oci_managed_kafka_kafka_cluster", ManagedKafkaKafkaClusterDataSource())
	tfresource.RegisterDatasource("oci_managed_kafka_kafka_cluster_config", ManagedKafkaKafkaClusterConfigDataSource())
	tfresource.RegisterDatasource("oci_managed_kafka_kafka_cluster_config_version", ManagedKafkaKafkaClusterConfigVersionDataSource())
	tfresource.RegisterDatasource("oci_managed_kafka_kafka_cluster_config_versions", ManagedKafkaKafkaClusterConfigVersionsDataSource())
	tfresource.RegisterDatasource("oci_managed_kafka_kafka_cluster_configs", ManagedKafkaKafkaClusterConfigsDataSource())
	tfresource.RegisterDatasource("oci_managed_kafka_kafka_clusters", ManagedKafkaKafkaClustersDataSource())
}
