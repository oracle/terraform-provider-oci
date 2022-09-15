// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package provider

import (
	tf_adm "github.com/oracle/terraform-provider-oci/internal/service/adm"
	tf_ai_anomaly_detection "github.com/oracle/terraform-provider-oci/internal/service/ai_anomaly_detection"
	tf_ai_vision "github.com/oracle/terraform-provider-oci/internal/service/ai_vision"
	tf_analytics "github.com/oracle/terraform-provider-oci/internal/service/analytics"
	tf_announcements_service "github.com/oracle/terraform-provider-oci/internal/service/announcements_service"
	tf_apigateway "github.com/oracle/terraform-provider-oci/internal/service/apigateway"
	tf_apm "github.com/oracle/terraform-provider-oci/internal/service/apm"
	tf_apm_config "github.com/oracle/terraform-provider-oci/internal/service/apm_config"
	tf_apm_synthetics "github.com/oracle/terraform-provider-oci/internal/service/apm_synthetics"
	tf_apm_traces "github.com/oracle/terraform-provider-oci/internal/service/apm_traces"
	tf_appmgmt_control "github.com/oracle/terraform-provider-oci/internal/service/appmgmt_control"
	tf_artifacts "github.com/oracle/terraform-provider-oci/internal/service/artifacts"
	tf_audit "github.com/oracle/terraform-provider-oci/internal/service/audit"
	tf_autoscaling "github.com/oracle/terraform-provider-oci/internal/service/autoscaling"
	tf_bastion "github.com/oracle/terraform-provider-oci/internal/service/bastion"
	tf_bds "github.com/oracle/terraform-provider-oci/internal/service/bds"
	tf_blockchain "github.com/oracle/terraform-provider-oci/internal/service/blockchain"
	tf_budget "github.com/oracle/terraform-provider-oci/internal/service/budget"
	tf_certificates_management "github.com/oracle/terraform-provider-oci/internal/service/certificates_management"
	tf_cloud_guard "github.com/oracle/terraform-provider-oci/internal/service/cloud_guard"
	tf_cloud_migrations "github.com/oracle/terraform-provider-oci/internal/service/cloud_migrations"
	tf_computeinstanceagent "github.com/oracle/terraform-provider-oci/internal/service/computeinstanceagent"
	tf_containerengine "github.com/oracle/terraform-provider-oci/internal/service/containerengine"
	tf_core "github.com/oracle/terraform-provider-oci/internal/service/core"
	tf_data_connectivity "github.com/oracle/terraform-provider-oci/internal/service/data_connectivity"
	tf_data_labeling_service "github.com/oracle/terraform-provider-oci/internal/service/data_labeling_service"
	tf_data_safe "github.com/oracle/terraform-provider-oci/internal/service/data_safe"
	tf_database "github.com/oracle/terraform-provider-oci/internal/service/database"
	tf_database_management "github.com/oracle/terraform-provider-oci/internal/service/database_management"
	tf_database_migration "github.com/oracle/terraform-provider-oci/internal/service/database_migration"
	tf_database_tools "github.com/oracle/terraform-provider-oci/internal/service/database_tools"
	tf_datacatalog "github.com/oracle/terraform-provider-oci/internal/service/datacatalog"
	tf_dataflow "github.com/oracle/terraform-provider-oci/internal/service/dataflow"
	tf_dataintegration "github.com/oracle/terraform-provider-oci/internal/service/dataintegration"
	tf_datascience "github.com/oracle/terraform-provider-oci/internal/service/datascience"
	tf_devops "github.com/oracle/terraform-provider-oci/internal/service/devops"
	tf_dns "github.com/oracle/terraform-provider-oci/internal/service/dns"
	tf_em_warehouse "github.com/oracle/terraform-provider-oci/internal/service/em_warehouse"
	tf_email "github.com/oracle/terraform-provider-oci/internal/service/email"
	tf_events "github.com/oracle/terraform-provider-oci/internal/service/events"
	tf_file_storage "github.com/oracle/terraform-provider-oci/internal/service/file_storage"
	tf_functions "github.com/oracle/terraform-provider-oci/internal/service/functions"
	tf_fusion_apps "github.com/oracle/terraform-provider-oci/internal/service/fusion_apps"
	tf_generic_artifacts_content "github.com/oracle/terraform-provider-oci/internal/service/generic_artifacts_content"
	tf_golden_gate "github.com/oracle/terraform-provider-oci/internal/service/golden_gate"
	tf_health_checks "github.com/oracle/terraform-provider-oci/internal/service/health_checks"
	tf_identity "github.com/oracle/terraform-provider-oci/internal/service/identity"
	tf_identity_data_plane "github.com/oracle/terraform-provider-oci/internal/service/identity_data_plane"
	tf_integration "github.com/oracle/terraform-provider-oci/internal/service/integration"
	tf_jms "github.com/oracle/terraform-provider-oci/internal/service/jms"
	tf_kms "github.com/oracle/terraform-provider-oci/internal/service/kms"
	tf_license_manager "github.com/oracle/terraform-provider-oci/internal/service/license_manager"
	tf_limits "github.com/oracle/terraform-provider-oci/internal/service/limits"
	tf_load_balancer "github.com/oracle/terraform-provider-oci/internal/service/load_balancer"
	tf_log_analytics "github.com/oracle/terraform-provider-oci/internal/service/log_analytics"
	tf_logging "github.com/oracle/terraform-provider-oci/internal/service/logging"
	tf_management_agent "github.com/oracle/terraform-provider-oci/internal/service/management_agent"
	tf_management_dashboard "github.com/oracle/terraform-provider-oci/internal/service/management_dashboard"
	tf_marketplace "github.com/oracle/terraform-provider-oci/internal/service/marketplace"
	tf_metering_computation "github.com/oracle/terraform-provider-oci/internal/service/metering_computation"
	tf_monitoring "github.com/oracle/terraform-provider-oci/internal/service/monitoring"
	tf_mysql "github.com/oracle/terraform-provider-oci/internal/service/mysql"
	tf_network_firewall "github.com/oracle/terraform-provider-oci/internal/service/network_firewall"
	tf_network_load_balancer "github.com/oracle/terraform-provider-oci/internal/service/network_load_balancer"
	tf_nosql "github.com/oracle/terraform-provider-oci/internal/service/nosql"
	tf_objectstorage "github.com/oracle/terraform-provider-oci/internal/service/objectstorage"
	tf_oce "github.com/oracle/terraform-provider-oci/internal/service/oce"
	tf_ocvp "github.com/oracle/terraform-provider-oci/internal/service/ocvp"
	tf_oda "github.com/oracle/terraform-provider-oci/internal/service/oda"
	tf_onesubscription "github.com/oracle/terraform-provider-oci/internal/service/onesubscription"
	tf_ons "github.com/oracle/terraform-provider-oci/internal/service/ons"
	tf_opensearch "github.com/oracle/terraform-provider-oci/internal/service/opensearch"
	tf_operator_access_control "github.com/oracle/terraform-provider-oci/internal/service/operator_access_control"
	tf_opsi "github.com/oracle/terraform-provider-oci/internal/service/opsi"
	tf_optimizer "github.com/oracle/terraform-provider-oci/internal/service/optimizer"
	tf_osmanagement "github.com/oracle/terraform-provider-oci/internal/service/osmanagement"
	tf_osp_gateway "github.com/oracle/terraform-provider-oci/internal/service/osp_gateway"
	tf_osub_billing_schedule "github.com/oracle/terraform-provider-oci/internal/service/osub_billing_schedule"
	tf_osub_organization_subscription "github.com/oracle/terraform-provider-oci/internal/service/osub_organization_subscription"
	tf_osub_subscription "github.com/oracle/terraform-provider-oci/internal/service/osub_subscription"
	tf_osub_usage "github.com/oracle/terraform-provider-oci/internal/service/osub_usage"
	tf_resourcemanager "github.com/oracle/terraform-provider-oci/internal/service/resourcemanager"
	tf_sch "github.com/oracle/terraform-provider-oci/internal/service/sch"
	tf_secrets "github.com/oracle/terraform-provider-oci/internal/service/secrets"
	tf_service_catalog "github.com/oracle/terraform-provider-oci/internal/service/service_catalog"
	tf_service_manager_proxy "github.com/oracle/terraform-provider-oci/internal/service/service_manager_proxy"
	tf_service_mesh "github.com/oracle/terraform-provider-oci/internal/service/service_mesh"
	tf_stack_monitoring "github.com/oracle/terraform-provider-oci/internal/service/stack_monitoring"
	tf_streaming "github.com/oracle/terraform-provider-oci/internal/service/streaming"
	tf_usage_proxy "github.com/oracle/terraform-provider-oci/internal/service/usage_proxy"
	tf_vault "github.com/oracle/terraform-provider-oci/internal/service/vault"
	tf_visual_builder "github.com/oracle/terraform-provider-oci/internal/service/visual_builder"
	tf_vn_monitoring "github.com/oracle/terraform-provider-oci/internal/service/vn_monitoring"
	tf_vulnerability_scanning "github.com/oracle/terraform-provider-oci/internal/service/vulnerability_scanning"
	tf_waa "github.com/oracle/terraform-provider-oci/internal/service/waa"
	tf_waas "github.com/oracle/terraform-provider-oci/internal/service/waas"
	tf_waf "github.com/oracle/terraform-provider-oci/internal/service/waf"
)

func init() {
	tf_adm.RegisterResource()
	tf_ai_anomaly_detection.RegisterResource()
	tf_ai_vision.RegisterResource()
	tf_analytics.RegisterResource()
	tf_announcements_service.RegisterResource()
	tf_apigateway.RegisterResource()
	tf_apm.RegisterResource()
	tf_apm_config.RegisterResource()
	tf_apm_synthetics.RegisterResource()
	tf_apm_traces.RegisterResource()
	tf_appmgmt_control.RegisterResource()
	tf_artifacts.RegisterResource()
	tf_audit.RegisterResource()
	tf_autoscaling.RegisterResource()
	tf_bastion.RegisterResource()
	tf_bds.RegisterResource()
	tf_blockchain.RegisterResource()
	tf_budget.RegisterResource()
	tf_certificates_management.RegisterResource()
	tf_cloud_guard.RegisterResource()
	tf_cloud_migrations.RegisterResource()
	tf_computeinstanceagent.RegisterResource()
	tf_containerengine.RegisterResource()
	tf_core.RegisterResource()
	tf_data_connectivity.RegisterResource()
	tf_data_labeling_service.RegisterResource()
	tf_data_safe.RegisterResource()
	tf_database.RegisterResource()
	tf_database_management.RegisterResource()
	tf_database_migration.RegisterResource()
	tf_database_tools.RegisterResource()
	tf_datacatalog.RegisterResource()
	tf_dataflow.RegisterResource()
	tf_dataintegration.RegisterResource()
	tf_datascience.RegisterResource()
	tf_devops.RegisterResource()
	tf_dns.RegisterResource()
	tf_em_warehouse.RegisterResource()
	tf_email.RegisterResource()
	tf_events.RegisterResource()
	tf_file_storage.RegisterResource()
	tf_functions.RegisterResource()
	tf_fusion_apps.RegisterResource()
	tf_generic_artifacts_content.RegisterResource()
	tf_golden_gate.RegisterResource()
	tf_health_checks.RegisterResource()
	tf_identity.RegisterResource()
	tf_identity_data_plane.RegisterResource()
	tf_integration.RegisterResource()
	tf_jms.RegisterResource()
	tf_kms.RegisterResource()
	tf_license_manager.RegisterResource()
	tf_limits.RegisterResource()
	tf_load_balancer.RegisterResource()
	tf_log_analytics.RegisterResource()
	tf_logging.RegisterResource()
	tf_management_agent.RegisterResource()
	tf_management_dashboard.RegisterResource()
	tf_marketplace.RegisterResource()
	tf_metering_computation.RegisterResource()
	tf_monitoring.RegisterResource()
	tf_mysql.RegisterResource()
	tf_network_firewall.RegisterResource()
	tf_network_load_balancer.RegisterResource()
	tf_nosql.RegisterResource()
	tf_objectstorage.RegisterResource()
	tf_oce.RegisterResource()
	tf_ocvp.RegisterResource()
	tf_oda.RegisterResource()
	tf_onesubscription.RegisterResource()
	tf_ons.RegisterResource()
	tf_opensearch.RegisterResource()
	tf_operator_access_control.RegisterResource()
	tf_opsi.RegisterResource()
	tf_optimizer.RegisterResource()
	tf_osmanagement.RegisterResource()
	tf_osp_gateway.RegisterResource()
	tf_osub_billing_schedule.RegisterResource()
	tf_osub_organization_subscription.RegisterResource()
	tf_osub_subscription.RegisterResource()
	tf_osub_usage.RegisterResource()
	tf_resourcemanager.RegisterResource()
	tf_sch.RegisterResource()
	tf_secrets.RegisterResource()
	tf_service_catalog.RegisterResource()
	tf_service_manager_proxy.RegisterResource()
	tf_service_mesh.RegisterResource()
	tf_stack_monitoring.RegisterResource()
	tf_streaming.RegisterResource()
	tf_usage_proxy.RegisterResource()
	tf_vault.RegisterResource()
	tf_visual_builder.RegisterResource()
	tf_vn_monitoring.RegisterResource()
	tf_vulnerability_scanning.RegisterResource()
	tf_waa.RegisterResource()
	tf_waas.RegisterResource()
	tf_waf.RegisterResource()
}
