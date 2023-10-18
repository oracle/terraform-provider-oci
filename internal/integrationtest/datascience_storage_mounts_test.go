package integrationtest

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	// Storage Mount (SM) Representation for network resources
	SMSubnetRepresentation = map[string]interface{}{
		"cidr_block":                 acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/24`},
		"compartment_id":             acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":                     acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"dns_label":                  acctest.Representation{RepType: acctest.Required, Create: `testsubnet`},
		"prohibit_public_ip_on_vnic": acctest.Representation{RepType: acctest.Required, Create: `true`},
		"lifecycle":                  acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRep},
		"security_list_ids":          acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_core_security_list.test_sec_list.id}`}},
		"route_table_id":             acctest.Representation{RepType: acctest.Required, Create: `${oci_core_default_route_table.default_route_table.id}`},
	}

	SMVcnRepresentation = map[string]interface{}{
		"cidr_block":     acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/16`},
		"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"dns_label":      acctest.Representation{RepType: acctest.Required, Create: `testvcn`},
		"lifecycle":      acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesRep},
	}

	SMSecurityListRepresentation = map[string]interface{}{
		"compartment_id":        acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"vcn_id":                acctest.Representation{RepType: acctest.Required, Create: `${oci_core_vcn.test_vcn.id}`},
		"egress_security_rules": []acctest.RepresentationGroup{{RepType: acctest.Required, Group: SMSecurityListTCPEgressSecurityRulesRepresentation}, {RepType: acctest.Required, Group: SMSecurityListServiceGatewayEgressSecurityRulesRepresentation}},
		"ingress_security_rules": []acctest.RepresentationGroup{{RepType: acctest.Required, Group: SMSecurityListSSHIngressSecurityRulesRepresentation}, {RepType: acctest.Required, Group: SMSecurityListICMPIngressSecurityRulesRepresentation}, {RepType: acctest.Required, Group: SMSecurityListICMPVcnCidrIngressSecurityRulesRepresentation},
			{RepType: acctest.Required, Group: SMSecurityListTCPIngressSecurityRulesRepresentation1}, {RepType: acctest.Required, Group: SMSecurityListTCPIngressSecurityRulesRepresentation2}, {RepType: acctest.Required, Group: SMSecurityListUDPIngressSecurityRulesRepresentation1}, {RepType: acctest.Required, Group: SMSecurityListUDPIngressSecurityRulesRepresentation2}},
	}

	SMSecurityListTCPEgressSecurityRulesRepresentation = map[string]interface{}{
		"destination": acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `all`},
		"stateless":   acctest.Representation{RepType: acctest.Required, Create: `false`},
	}

	SMSecurityListServiceGatewayEgressSecurityRulesRepresentation = map[string]interface{}{
		"destination":      acctest.Representation{RepType: acctest.Required, Create: `${lookup(data.oci_core_services.test_services.services[1], "cidr_block")}`},
		"destination_type": acctest.Representation{RepType: acctest.Required, Create: `SERVICE_CIDR_BLOCK`},
		"protocol":         acctest.Representation{RepType: acctest.Required, Create: `all`},
		"stateless":        acctest.Representation{RepType: acctest.Required, Create: `false`},
	}

	SMSecurityListSSHIngressSecurityRulesRepresentation = map[string]interface{}{
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `6`},
		"source":      acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"stateless":   acctest.Representation{RepType: acctest.Required, Create: `false`},
		"tcp_options": acctest.RepresentationGroup{RepType: acctest.Required, Group: SMSecurityListSSHIngressSecurityRulesTcpOptionsRepresentation},
	}

	SMSecurityListSSHIngressSecurityRulesTcpOptionsRepresentation = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `22`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `22`},
	}

	SMSecurityListICMPIngressSecurityRulesRepresentation = map[string]interface{}{
		"protocol":     acctest.Representation{RepType: acctest.Required, Create: `1`},
		"source":       acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"stateless":    acctest.Representation{RepType: acctest.Required, Create: `false`},
		"icmp_options": acctest.RepresentationGroup{RepType: acctest.Required, Group: SMSecurityListICMPIngressSecurityRulesTcpOptionsRepresentation1},
	}

	SMSecurityListICMPIngressSecurityRulesTcpOptionsRepresentation1 = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `3`},
		"code": acctest.Representation{RepType: acctest.Required, Create: `4`},
	}

	SMSecurityListICMPVcnCidrIngressSecurityRulesRepresentation = map[string]interface{}{
		"protocol":     acctest.Representation{RepType: acctest.Required, Create: `1`},
		"source":       acctest.Representation{RepType: acctest.Required, Create: `10.0.0.0/16`},
		"stateless":    acctest.Representation{RepType: acctest.Required, Create: `false`},
		"icmp_options": acctest.RepresentationGroup{RepType: acctest.Required, Group: SMSecurityListICMPIngressSecurityRulesTcpOptionsRepresentation2},
	}

	SMSecurityListICMPIngressSecurityRulesTcpOptionsRepresentation2 = map[string]interface{}{
		"type": acctest.Representation{RepType: acctest.Required, Create: `3`},
	}

	SMSecurityListTCPIngressSecurityRulesRepresentation1 = map[string]interface{}{
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `6`},
		"source":      acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"stateless":   acctest.Representation{RepType: acctest.Required, Create: `false`},
		"tcp_options": acctest.RepresentationGroup{RepType: acctest.Required, Group: SMSecurityListTCPIngressSecurityRulesTcpOptionsRepresentation1},
	}

	SMSecurityListTCPIngressSecurityRulesTcpOptionsRepresentation1 = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `111`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `111`},
	}

	SMSecurityListTCPIngressSecurityRulesRepresentation2 = map[string]interface{}{
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `6`},
		"source":      acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"stateless":   acctest.Representation{RepType: acctest.Required, Create: `false`},
		"tcp_options": acctest.RepresentationGroup{RepType: acctest.Required, Group: SMSecurityListTCPIngressSecurityRulesTcpOptionsRepresentation2},
	}

	SMSecurityListTCPIngressSecurityRulesTcpOptionsRepresentation2 = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `2050`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `2048`},
	}

	SMSecurityListUDPIngressSecurityRulesRepresentation1 = map[string]interface{}{
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `17`},
		"source":      acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"stateless":   acctest.Representation{RepType: acctest.Required, Create: `false`},
		"udp_options": acctest.RepresentationGroup{RepType: acctest.Required, Group: SMSecurityListUDPIngressSecurityRulesUdpOptionsRepresentation1},
	}

	SMSecurityListUDPIngressSecurityRulesUdpOptionsRepresentation1 = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `111`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `111`},
	}

	SMSecurityListUDPIngressSecurityRulesRepresentation2 = map[string]interface{}{
		"protocol":    acctest.Representation{RepType: acctest.Required, Create: `17`},
		"source":      acctest.Representation{RepType: acctest.Required, Create: `0.0.0.0/0`},
		"stateless":   acctest.Representation{RepType: acctest.Required, Create: `false`},
		"udp_options": acctest.RepresentationGroup{RepType: acctest.Required, Group: SMSecurityListUDPIngressSecurityRulesUdpOptionsRepresentation2},
	}

	SMSecurityListUDPIngressSecurityRulesUdpOptionsRepresentation2 = map[string]interface{}{
		"max": acctest.Representation{RepType: acctest.Required, Create: `2050`},
		"min": acctest.Representation{RepType: acctest.Required, Create: `2048`},
	}

	// Storage Mount Representation
	StorageMountConfigurationDetailsListRepresentation = map[string]interface{}{
		"destination_directory_name": acctest.Representation{RepType: acctest.Required, Create: `destinationDirectoryName`, Update: `destinationDirectoryName2`},
		"destination_path":           acctest.Representation{RepType: acctest.Optional, Create: `/destinationPath`, Update: `/destinationPath2`},
		"storage_type":               acctest.Representation{RepType: acctest.Required, Create: `FILE_STORAGE`, Update: `OBJECT_STORAGE`},
		"bucket":                     acctest.Representation{RepType: acctest.Optional, Create: ``, Update: `${oci_objectstorage_bucket.test_bucket.name}`},
		"export_id":                  acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_export.test_export.id}`, Update: ``},
		"mount_target_id":            acctest.Representation{RepType: acctest.Optional, Create: `${oci_file_storage_mount_target.test_mount_target.id}`, Update: ``},
		"namespace":                  acctest.Representation{RepType: acctest.Optional, Create: ``, Update: `${data.oci_objectstorage_namespace.test_namespace.namespace}`},
		"prefix":                     acctest.Representation{RepType: acctest.Optional, Create: ``, Update: `prefix`},
	}

	// Storage Mount Dependencies
	StorageMountConfigurationDetailsListResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, SMSubnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, SMVcnRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_security_list", "test_sec_list", acctest.Required, acctest.Create, SMSecurityListRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_default_route_table", "default_route_table", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(routeTablesRepresentation, map[string]interface{}{
				"route_rules": acctest.RepresentationGroup{RepType: acctest.Required, Group: allOCIServiceRouteTableRouteRulesRepresentationWithServiceCidr},
			})) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_core_services", "test_services", acctest.Required, acctest.Create, CoreCoreServiceDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_service_gateway", "test_service_gateway", acctest.Required, acctest.Create,
			acctest.RepresentationCopyWithNewProperties(CoreServiceGatewayRepresentation, map[string]interface{}{
				"services": acctest.RepresentationGroup{RepType: acctest.Required, Group: allOCIServiceGatewayServicesRepresentation},
			})) +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_export_set", "test_export_set", acctest.Optional, acctest.Create, FileStorageExportSetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_export", "test_export", acctest.Optional, acctest.Create, FileStorageExportRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_file_system", "test_file_system", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithRemovedProperties(FileStorageFileSystemRepresentation, []string{"filesystem_snapshot_policy_id", "kms_key_id", "source_snapshot_id"})) +
		acctest.GenerateResourceFromRepresentationMap("oci_file_storage_mount_target", "test_mount_target", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithRemovedProperties(FileStorageMountTargetRepresentation, []string{"defined_tags", "kerberos", "ldap_idmap", "nsg_ids", "ip_address", "idmap_type", "hostname_label"})) +
		acctest.GenerateResourceFromRepresentationMap("oci_objectstorage_bucket", "test_bucket", acctest.Optional, acctest.Create, acctest.RepresentationCopyWithRemovedProperties(ObjectStorageBucketRepresentation, []string{"kms_key_id", "defined_tags"})) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_objectstorage_namespace", "test_namespace", acctest.Optional, acctest.Create, ObjectStorageObjectStorageNamespaceSingularDataSourceRepresentation)

	// Job Storage Mounts
	JobResourceStorageMountConfigurationDetailsListConfig = JobResourceStorageMountConfigurationDetailsListResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_job", "test_job", acctest.Optional, acctest.Update, JobResourceStorageMountConfigurationDetailsListRepresentation)

	jobStorageMountConfigurationDetailsListSingularDataSourceRepresentation = map[string]interface{}{
		"job_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_job.test_job.id}`},
	}

	JobResourceStorageMountConfigurationDetailsListRepresentation = map[string]interface{}{
		"compartment_id":                               acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"job_configuration_details":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceJobJobConfigurationDetailsRepresentation},
		"job_infrastructure_configuration_details":     acctest.RepresentationGroup{RepType: acctest.Required, Group: DatascienceJobJobInfrastructureConfigurationDetailsRepresentation},
		"job_storage_mount_configuration_details_list": acctest.RepresentationGroup{RepType: acctest.Optional, Group: StorageMountConfigurationDetailsListRepresentation},
		"project_id":                   acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"job_artifact":                 acctest.Representation{RepType: acctest.Optional, Create: `../../examples/datascience/job-artifact.py`},
		"artifact_content_length":      acctest.Representation{RepType: acctest.Optional, Create: `1380`}, // wc -c job-artifact.py
		"artifact_content_disposition": acctest.Representation{RepType: acctest.Optional, Create: `attachment; filename=job-artifact.py`},
		"defined_tags":                 acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":                  acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":                 acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"delete_related_job_runs":      acctest.Representation{RepType: acctest.Optional, Create: `false`, Update: `true`},
		"lifecycle":                    acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreMlJobDefinedTagsChangesRepresentation},
	}

	JobResourceStorageMountConfigurationDetailsListResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		acctest.GenerateDataSourceFromRepresentationMap("oci_core_shapes", "test_shapes", acctest.Required, acctest.Create, CoreCoreShapeDataSourceRepresentation) +
		StorageMountConfigurationDetailsListResourceDependencies +
		AvailabilityDomainConfig +
		DefinedTagsDependencies

	// Notebook Session Storage Mounts
	NotebookSessionStorageMountConfigurationDetailsListResourceConfig = NotebookSessionStorageMountConfigurationDetailsListResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Update, NotebookSessionStorageMountConfigurationDetailsListRepresentation)

	notebookSessionStorageMountConfigurationDetailsListSingularDataSourceRepresentation = map[string]interface{}{
		"notebook_session_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_notebook_session.test_notebook_session.id}`},
	}

	notebookSessionWithStorageMountsConfigDetailsRepresentation = map[string]interface{}{
		"shape":                     acctest.Representation{RepType: acctest.Required, Create: `VM.Standard2.1`},
		"block_storage_size_in_gbs": acctest.Representation{RepType: acctest.Optional, Create: `100`},
		"subnet_id":                 acctest.Representation{RepType: acctest.Required, Create: `${oci_core_subnet.test_subnet.id}`},
		"private_endpoint_id":       acctest.Representation{RepType: acctest.Optional, Create: `${oci_datascience_private_endpoint.test_data_science_private_endpoint.id}`},
	}

	NotebookSessionStorageMountConfigurationDetailsListRepresentation = map[string]interface{}{
		"compartment_id":                  acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"project_id":                      acctest.Representation{RepType: acctest.Required, Create: `${oci_datascience_project.test_project.id}`},
		"defined_tags":                    acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"display_name":                    acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":                   acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"Department": "Finance"}, Update: map[string]string{"Department": "Accounting"}},
		"notebook_session_config_details": acctest.RepresentationGroup{RepType: acctest.Required, Group: notebookSessionWithStorageMountsConfigDetailsRepresentation},
		"notebook_session_storage_mount_configuration_details_list": acctest.RepresentationGroup{RepType: acctest.Optional, Group: StorageMountConfigurationDetailsListRepresentation},
		"lifecycle": acctest.RepresentationGroup{RepType: acctest.Optional, Group: definedTagsIgnoreRepresentation},
	}

	NotebookSessionStorageMountConfigurationDetailsListResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_datascience_project", "test_project", acctest.Required, acctest.Create, DatascienceProjectRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_private_endpoint", "test_data_science_private_endpoint", acctest.Required, acctest.Create, DataSciencePrivateEndpointRepresentation) +
		StorageMountConfigurationDetailsListResourceDependencies +
		AvailabilityDomainConfig +
		DefinedTagsDependencies
)

func TestDatascienceJobResourceWithStorageMountConfigurationDetailsList_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceJobResourceWithStorageMountConfigurationDetailsList_basic")
	defer httpreplay.SaveScenario()

	provider := acctest.TestAccProvider
	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_job.test_job"
	singularDatasourceName := "data.oci_datascience_job.test_job"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+JobResourceStorageMountConfigurationDetailsListResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_job", "test_job", acctest.Optional, acctest.Create, JobResourceStorageMountConfigurationDetailsListRepresentation), "datascience", "job", t)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acctest.TestAccPreCheck(t) },
		Providers: map[string]*schema.Provider{
			"oci": provider,
		},
		CheckDestroy: testAccCheckDatascienceJobDestroy,
		Steps: []resource.TestStep{
			// verify Create
			{
				Config: config + compartmentIdVariableStr + JobResourceStorageMountConfigurationDetailsListResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_datascience_job", "test_job", acctest.Required, acctest.Create, JobResourceStorageMountConfigurationDetailsListRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.job_type", "DEFAULT"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.0.block_storage_size_in_gbs", "50"),
					resource.TestCheckResourceAttrSet(resourceName, "job_infrastructure_configuration_details.0.shape_name"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						return err
					},
				),
			},
			// delete before next Create
			{
				Config: config + compartmentIdVariableStr + JobResourceStorageMountConfigurationDetailsListResourceDependencies,
			},
			// verify Create with optionals
			{
				Config: config + compartmentIdVariableStr + JobResourceStorageMountConfigurationDetailsListResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_datascience_job", "test_job", acctest.Optional, acctest.Create, JobResourceStorageMountConfigurationDetailsListRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.command_line_arguments", "commandLineArguments"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.environment_variables.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.job_type", "DEFAULT"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.maximum_runtime_in_minutes", "10"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.0.block_storage_size_in_gbs", "50"),
					resource.TestCheckResourceAttrSet(resourceName, "job_infrastructure_configuration_details.0.shape_name"),
					resource.TestCheckResourceAttr(resourceName, "job_storage_mount_configuration_details_list.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_storage_mount_configuration_details_list.0.destination_directory_name", "destinationDirectoryName"),
					resource.TestCheckResourceAttr(resourceName, "job_storage_mount_configuration_details_list.0.destination_path", "/destinationPath"),
					resource.TestCheckResourceAttrSet(resourceName, "job_storage_mount_configuration_details_list.0.export_id"),
					resource.TestCheckResourceAttrSet(resourceName, "job_storage_mount_configuration_details_list.0.mount_target_id"),
					resource.TestCheckResourceAttr(resourceName, "job_storage_mount_configuration_details_list.0.storage_type", "FILE_STORAGE"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId, err = acctest.FromInstanceState(s, resourceName, "id")
						if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
							if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
								return errExport
							}
						}
						return err
					},
				),
			},

			// verify Update to the compartment (the compartment will be switched back in the next step)
			{
				Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + JobResourceStorageMountConfigurationDetailsListResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_datascience_job", "test_job", acctest.Optional, acctest.Create,
						acctest.RepresentationCopyWithNewProperties(JobResourceStorageMountConfigurationDetailsListRepresentation, map[string]interface{}{
							"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
						})),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentIdU),
					resource.TestCheckResourceAttr(resourceName, "description", "description"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.command_line_arguments", "commandLineArguments"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.environment_variables.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.job_type", "DEFAULT"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.maximum_runtime_in_minutes", "10"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.0.block_storage_size_in_gbs", "50"),
					resource.TestCheckResourceAttr(resourceName, "job_storage_mount_configuration_details_list.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_storage_mount_configuration_details_list.0.destination_directory_name", "destinationDirectoryName"),
					resource.TestCheckResourceAttr(resourceName, "job_storage_mount_configuration_details_list.0.destination_path", "/destinationPath"),
					resource.TestCheckResourceAttrSet(resourceName, "job_storage_mount_configuration_details_list.0.export_id"),
					resource.TestCheckResourceAttrSet(resourceName, "job_storage_mount_configuration_details_list.0.mount_target_id"),
					resource.TestCheckResourceAttr(resourceName, "job_storage_mount_configuration_details_list.0.storage_type", "FILE_STORAGE"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("resource recreated when it was supposed to be updated")
						}
						return err
					},
				),
			},

			// verify updates to updatable parameters
			{
				Config: config + compartmentIdVariableStr + JobResourceStorageMountConfigurationDetailsListResourceDependencies +
					acctest.GenerateResourceFromRepresentationMap("oci_datascience_job", "test_job", acctest.Optional, acctest.Update, JobResourceStorageMountConfigurationDetailsListRepresentation),
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(resourceName, "description", "description2"),
					resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.command_line_arguments", "commandLineArguments"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.environment_variables.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.job_type", "DEFAULT"),
					resource.TestCheckResourceAttr(resourceName, "job_configuration_details.0.maximum_runtime_in_minutes", "10"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "job_infrastructure_configuration_details.0.block_storage_size_in_gbs", "50"),
					resource.TestCheckResourceAttr(resourceName, "job_storage_mount_configuration_details_list.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "job_storage_mount_configuration_details_list.0.bucket"),
					resource.TestCheckResourceAttr(resourceName, "job_storage_mount_configuration_details_list.0.destination_directory_name", "destinationDirectoryName2"),
					resource.TestCheckResourceAttr(resourceName, "job_storage_mount_configuration_details_list.0.destination_path", "/destinationPath2"),
					resource.TestCheckResourceAttr(resourceName, "job_storage_mount_configuration_details_list.0.prefix", "prefix"),
					resource.TestCheckResourceAttr(resourceName, "job_storage_mount_configuration_details_list.0.storage_type", "OBJECT_STORAGE"),
					resource.TestCheckResourceAttrSet(resourceName, "job_infrastructure_configuration_details.0.shape_name"),
					resource.TestCheckResourceAttrSet(resourceName, "project_id"),
					resource.TestCheckResourceAttrSet(resourceName, "state"),
					resource.TestCheckResourceAttrSet(resourceName, "time_created"),

					func(s *terraform.State) (err error) {
						resId2, err = acctest.FromInstanceState(s, resourceName, "id")
						if resId != resId2 {
							return fmt.Errorf("Resource recreated when it was supposed to be updated.")
						}
						return err
					},
				),
			},
			// verify singular datasource
			{
				Config: config +
					acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_job", "test_job", acctest.Required, acctest.Create, jobStorageMountConfigurationDetailsListSingularDataSourceRepresentation) +
					compartmentIdVariableStr + JobResourceStorageMountConfigurationDetailsListConfig,
				Check: acctest.ComposeAggregateTestCheckFuncWrapper(
					resource.TestCheckResourceAttrSet(singularDatasourceName, "job_id"),

					resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
					resource.TestCheckResourceAttr(singularDatasourceName, "description", "description2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
					resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_configuration_details.0.command_line_arguments", "commandLineArguments"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_configuration_details.0.environment_variables.%", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_configuration_details.0.job_type", "DEFAULT"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_configuration_details.0.maximum_runtime_in_minutes", "10"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_infrastructure_configuration_details.#", "1"),
					resource.TestCheckResourceAttr(singularDatasourceName, "job_infrastructure_configuration_details.0.block_storage_size_in_gbs", "50"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "job_storage_mount_configuration_details_list.0.namespace"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
					resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				),
			},
			// verify resource import
			{
				Config:            config + DatascienceJobRequiredOnlyResource,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"artifact_content_disposition",
					"artifact_content_length",
					"lifecycle_details",
					"delete_related_job_runs",
					"job_artifact",
				},
				ResourceName: resourceName,
			},
		},
	})
}

func TestDatascienceNotebookSessionResourceWithStorageMountConfigurationDetailsList_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatascienceNotebookSessionResourceWithStorageMountConfigurationDetailsList_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	compartmentIdU := utils.GetEnvSettingWithDefault("compartment_id_for_update", compartmentId)
	compartmentIdUVariableStr := fmt.Sprintf("variable \"compartment_id_for_update\" { default = \"%s\" }\n", compartmentIdU)

	resourceName := "oci_datascience_notebook_session.test_notebook_session"
	singularDatasourceName := "data.oci_datascience_notebook_session.test_notebook_session"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+NotebookSessionStorageMountConfigurationDetailsListResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Create, NotebookSessionStorageMountConfigurationDetailsListRepresentation), "datascience", "notebookSession", t)

	acctest.ResourceTest(t, testAccCheckDatascienceNotebookSessionDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + NotebookSessionStorageMountConfigurationDetailsListResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Required, acctest.Create, NotebookSessionStorageMountConfigurationDetailsListRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),

				func(s *terraform.State) (err error) {
					_, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + NotebookSessionStorageMountConfigurationDetailsListResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + NotebookSessionStorageMountConfigurationDetailsListResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Create, NotebookSessionStorageMountConfigurationDetailsListRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "notebook_session_config_details.0.private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_storage_mount_configuration_details_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_storage_mount_configuration_details_list.0.destination_directory_name", "destinationDirectoryName"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_storage_mount_configuration_details_list.0.destination_path", "/destinationPath"),
				resource.TestCheckResourceAttrSet(resourceName, "notebook_session_storage_mount_configuration_details_list.0.export_id"),
				resource.TestCheckResourceAttrSet(resourceName, "notebook_session_storage_mount_configuration_details_list.0.mount_target_id"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_storage_mount_configuration_details_list.0.storage_type", "FILE_STORAGE"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},
		// verify Update to the compartment (the compartment will be switched back in the next step)
		{
			Config: config + compartmentIdVariableStr + compartmentIdUVariableStr + NotebookSessionStorageMountConfigurationDetailsListResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Create,
					acctest.RepresentationCopyWithNewProperties(NotebookSessionStorageMountConfigurationDetailsListRepresentation, map[string]interface{}{
						"compartment_id": acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id_for_update}`},
					})),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "notebook_session_config_details.0.private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_storage_mount_configuration_details_list.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_storage_mount_configuration_details_list.0.destination_directory_name", "destinationDirectoryName"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_storage_mount_configuration_details_list.0.destination_path", "/destinationPath"),
				resource.TestCheckResourceAttrSet(resourceName, "notebook_session_storage_mount_configuration_details_list.0.export_id"),
				resource.TestCheckResourceAttrSet(resourceName, "notebook_session_storage_mount_configuration_details_list.0.mount_target_id"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_storage_mount_configuration_details_list.0.storage_type", "FILE_STORAGE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("resource recreated when it was supposed to be updated")
					}
					return err
				},
			),
		},
		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + NotebookSessionStorageMountConfigurationDetailsListResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Optional, acctest.Update, NotebookSessionStorageMountConfigurationDetailsListRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "created_by"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttrSet(resourceName, "project_id"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "notebook_session_config_details.0.private_endpoint_id"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_storage_mount_configuration_details_list.#", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "notebook_session_storage_mount_configuration_details_list.0.bucket"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_storage_mount_configuration_details_list.0.destination_directory_name", "destinationDirectoryName2"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_storage_mount_configuration_details_list.0.destination_path", "/destinationPath2"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_storage_mount_configuration_details_list.0.prefix", "prefix"),
				resource.TestCheckResourceAttr(resourceName, "notebook_session_storage_mount_configuration_details_list.0.storage_type", "OBJECT_STORAGE"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_datascience_notebook_session", "test_notebook_session", acctest.Required, acctest.Create, notebookSessionStorageMountConfigurationDetailsListSingularDataSourceRepresentation) +
				compartmentIdVariableStr + NotebookSessionStorageMountConfigurationDetailsListResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "notebook_session_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "created_by"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(singularDatasourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "notebook_session_url"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "notebook_session_storage_mount_configuration_details_list.0.namespace"),
			),
		},
		// verify resource import
		{
			Config:                  config + DatascienceNotebookSessionRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}
