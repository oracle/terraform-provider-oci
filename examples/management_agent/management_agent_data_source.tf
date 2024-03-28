// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

resource "oci_management_agent_management_agent_data_source" "add_datasource" {
  compartment_id      = var.compartment_ocid
  management_agent_id = data.oci_management_agent_management_agents.find_agent.management_agents[0].id
  name                = "PrometheusTestA"
  type                = "PROMETHEUS_EMITTER"
  url                 = "http://localhost:1234"
  namespace           = "namespace"
  allow_metrics       = "*"

  depends_on = [oci_management_agent_management_agent.test_management_agent]
}
data "oci_management_agent_management_agent_data_sources" "test_datasource_data" {
  management_agent_id = data.oci_management_agent_management_agents.find_agent.management_agents[0].id
  name = "PrometheusTestA"
}
