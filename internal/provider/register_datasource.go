// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package provider

import (
	tf_adm "github.com/oracle/terraform-provider-oci/internal/service/adm"
	tf_ai_anomaly_detection "github.com/oracle/terraform-provider-oci/internal/service/ai_anomaly_detection"
	tf_ai_document "github.com/oracle/terraform-provider-oci/internal/service/ai_document"
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
	tf_certificates "github.com/oracle/terraform-provider-oci/internal/service/certificates"
	tf_certificates_management "github.com/oracle/terraform-provider-oci/internal/service/certificates_management"
	tf_cloud_bridge "github.com/oracle/terraform-provider-oci/internal/service/cloud_bridge"
	tf_cloud_guard "github.com/oracle/terraform-provider-oci/internal/service/cloud_guard"
	tf_cloud_migrations "github.com/oracle/terraform-provider-oci/internal/service/cloud_migrations"
	tf_computeinstanceagent "github.com/oracle/terraform-provider-oci/internal/service/computeinstanceagent"
	tf_container_instances "github.com/oracle/terraform-provider-oci/internal/service/container_instances"
	tf_containerengine "github.com/oracle/terraform-provider-oci/internal/service/containerengine"
	tf_core "github.com/oracle/terraform-provider-oci/internal/service/core"
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
	tf_disaster_recovery "github.com/oracle/terraform-provider-oci/internal/service/disaster_recovery"
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
	tf_identity_domains "github.com/oracle/terraform-provider-oci/internal/service/identity_domains"
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
	tf_media_services "github.com/oracle/terraform-provider-oci/internal/service/media_services"
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
	tf_opa "github.com/oracle/terraform-provider-oci/internal/service/opa"
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
	tf_queue "github.com/oracle/terraform-provider-oci/internal/service/queue"
	tf_recovery "github.com/oracle/terraform-provider-oci/internal/service/recovery"
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
	tf_vbs_inst "github.com/oracle/terraform-provider-oci/internal/service/vbs_inst"
	tf_visual_builder "github.com/oracle/terraform-provider-oci/internal/service/visual_builder"
	tf_vn_monitoring "github.com/oracle/terraform-provider-oci/internal/service/vn_monitoring"
	tf_vulnerability_scanning "github.com/oracle/terraform-provider-oci/internal/service/vulnerability_scanning"
	tf_waa "github.com/oracle/terraform-provider-oci/internal/service/waa"
	tf_waas "github.com/oracle/terraform-provider-oci/internal/service/waas"
	tf_waf "github.com/oracle/terraform-provider-oci/internal/service/waf"
)

func init() {
	tf_adm.RegisterDatasource()
	tf_ai_anomaly_detection.RegisterDatasource()
	tf_ai_document.RegisterDatasource()
	tf_ai_vision.RegisterDatasource()
	tf_analytics.RegisterDatasource()
	tf_announcements_service.RegisterDatasource()
	tf_apigateway.RegisterDatasource()
	tf_apm.RegisterDatasource()
	tf_apm_config.RegisterDatasource()
	tf_apm_synthetics.RegisterDatasource()
	tf_apm_traces.RegisterDatasource()
	tf_appmgmt_control.RegisterDatasource()
	tf_artifacts.RegisterDatasource()
	tf_audit.RegisterDatasource()
	tf_autoscaling.RegisterDatasource()
	tf_bastion.RegisterDatasource()
	tf_bds.RegisterDatasource()
	tf_blockchain.RegisterDatasource()
	tf_budget.RegisterDatasource()
	tf_certificates.RegisterDatasource()
	tf_certificates_management.RegisterDatasource()
	tf_cloud_bridge.RegisterDatasource()
	tf_cloud_guard.RegisterDatasource()
	tf_cloud_migrations.RegisterDatasource()
	tf_computeinstanceagent.RegisterDatasource()
	tf_container_instances.RegisterDatasource()
	tf_containerengine.RegisterDatasource()
	tf_core.RegisterDatasource()
	tf_data_labeling_service.RegisterDatasource()
	tf_data_safe.RegisterDatasource()
	tf_database.RegisterDatasource()
	tf_database_management.RegisterDatasource()
	tf_database_migration.RegisterDatasource()
	tf_database_tools.RegisterDatasource()
	tf_datacatalog.RegisterDatasource()
	tf_dataflow.RegisterDatasource()
	tf_dataintegration.RegisterDatasource()
	tf_datascience.RegisterDatasource()
	tf_devops.RegisterDatasource()
	tf_disaster_recovery.RegisterDatasource()
	tf_dns.RegisterDatasource()
	tf_em_warehouse.RegisterDatasource()
	tf_email.RegisterDatasource()
	tf_events.RegisterDatasource()
	tf_file_storage.RegisterDatasource()
	tf_functions.RegisterDatasource()
	tf_fusion_apps.RegisterDatasource()
	tf_generic_artifacts_content.RegisterDatasource()
	tf_golden_gate.RegisterDatasource()
	tf_health_checks.RegisterDatasource()
	tf_identity.RegisterDatasource()
	tf_identity_data_plane.RegisterDatasource()
	tf_identity_domains.RegisterDatasource()
	tf_integration.RegisterDatasource()
	tf_jms.RegisterDatasource()
	tf_kms.RegisterDatasource()
	tf_license_manager.RegisterDatasource()
	tf_limits.RegisterDatasource()
	tf_load_balancer.RegisterDatasource()
	tf_log_analytics.RegisterDatasource()
	tf_logging.RegisterDatasource()
	tf_management_agent.RegisterDatasource()
	tf_management_dashboard.RegisterDatasource()
	tf_marketplace.RegisterDatasource()
	tf_media_services.RegisterDatasource()
	tf_metering_computation.RegisterDatasource()
	tf_monitoring.RegisterDatasource()
	tf_mysql.RegisterDatasource()
	tf_network_firewall.RegisterDatasource()
	tf_network_load_balancer.RegisterDatasource()
	tf_nosql.RegisterDatasource()
	tf_objectstorage.RegisterDatasource()
	tf_oce.RegisterDatasource()
	tf_ocvp.RegisterDatasource()
	tf_oda.RegisterDatasource()
	tf_onesubscription.RegisterDatasource()
	tf_ons.RegisterDatasource()
	tf_opa.RegisterDatasource()
	tf_opensearch.RegisterDatasource()
	tf_operator_access_control.RegisterDatasource()
	tf_opsi.RegisterDatasource()
	tf_optimizer.RegisterDatasource()
	tf_osmanagement.RegisterDatasource()
	tf_osp_gateway.RegisterDatasource()
	tf_osub_billing_schedule.RegisterDatasource()
	tf_osub_organization_subscription.RegisterDatasource()
	tf_osub_subscription.RegisterDatasource()
	tf_osub_usage.RegisterDatasource()
	tf_queue.RegisterDatasource()
	tf_recovery.RegisterDatasource()
	tf_resourcemanager.RegisterDatasource()
	tf_sch.RegisterDatasource()
	tf_secrets.RegisterDatasource()
	tf_service_catalog.RegisterDatasource()
	tf_service_manager_proxy.RegisterDatasource()
	tf_service_mesh.RegisterDatasource()
	tf_stack_monitoring.RegisterDatasource()
	tf_streaming.RegisterDatasource()
	tf_usage_proxy.RegisterDatasource()
	tf_vault.RegisterDatasource()
	tf_vbs_inst.RegisterDatasource()
	tf_visual_builder.RegisterDatasource()
	tf_vn_monitoring.RegisterDatasource()
	tf_vulnerability_scanning.RegisterDatasource()
	tf_waa.RegisterDatasource()
	tf_waas.RegisterDatasource()
	tf_waf.RegisterDatasource()
}
