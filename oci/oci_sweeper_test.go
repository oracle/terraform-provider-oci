package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func init() {
	// Calls to all sweeper init methods to initialize terraform sweeper framework with the sweepers

	initCoreAppCatalogSubscriptionSweeper()
	initDatabaseAutonomousDataWarehouseSweeper()
	initDatabaseAutonomousDatabaseSweeper()
	initLoadBalancerBackendSweeper()
	initLoadBalancerBackendSetSweeper()
	initDatabaseBackupSweeper()
	initCoreBootVolumeSweeper()
	initCoreBootVolumeBackupSweeper()
	initObjectStorageBucketSweeper()
	initLoadBalancerCertificateSweeper()
	initContainerengineClusterSweeper()
	initCoreConsoleHistorySweeper()
	initCoreCpeSweeper()
	initCoreCrossConnectSweeper()
	initCoreCrossConnectGroupSweeper()
	initDatabaseDbHomeSweeper()
	initCoreDhcpOptionsSweeper()
	initCoreDrgSweeper()
	initCoreDrgAttachmentSweeper()
	initFileStorageExportSweeper()
	initFileStorageFileSystemSweeper()
	initLoadBalancerHostnameSweeper()
	initCoreImageSweeper()
	initCoreInstanceSweeper()
	initCoreInstanceConfigurationSweeper()
	initCoreInstanceConsoleConnectionSweeper()
	initCoreInstancePoolSweeper()
	initCoreInternetGatewaySweeper()
	initCoreIpSecConnectionSweeper()
	initLoadBalancerLoadBalancerSweeper()
	initCoreLocalPeeringGatewaySweeper()
	initFileStorageMountTargetSweeper()
	initCoreNatGatewaySweeper()
	initContainerengineNodePoolSweeper()
	initObjectStorageObjectSweeper()
	initObjectStorageObjectLifecyclePolicySweeper()
	initLoadBalancerPathRouteSetSweeper()
	initObjectStoragePreauthenticatedRequestSweeper()
	initCorePrivateIpSweeper()
	initCorePublicIpSweeper()
	initCoreRemotePeeringConnectionSweeper()
	initCoreRouteTableSweeper()
	initCoreSecurityListSweeper()
	initEmailSenderSweeper()
	initCoreServiceGatewaySweeper()
	initFileStorageSnapshotSweeper()
	initCoreSubnetSweeper()
	initEmailSuppressionSweeper()
	initCoreVcnSweeper()
	initCoreVirtualCircuitSweeper()
	initCoreVnicAttachmentSweeper()
	initCoreVolumeSweeper()
	initCoreVolumeAttachmentSweeper()
	initCoreVolumeBackupSweeper()
	initCoreVolumeBackupPolicyAssignmentSweeper()
	initCoreVolumeGroupSweeper()
	initCoreVolumeGroupBackupSweeper()
	initDnsZoneSweeper()
}
