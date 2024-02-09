// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func createAdbInRegion(clients *tf_client.OracleClients, region string) (string, string, error) {
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")

	databaseClient, err := createDatabaseClient(clients, region)
	if err != nil {
		return "", "", err
	}

	cpuCoreCount := 1
	dataStorageSizeInTbs := 1
	adminPassword := string("BEstrO0ng_#11")
	dbName := utils.RandomString(1, utils.CharsetWithoutDigits) + utils.RandomString(13, utils.Charset)
	displayName := string("example_cr_source")
	dbVersion := "19c"
	isDedicated := false

	createAdbResponse, err := databaseClient.CreateAutonomousDatabase(context.Background(), oci_database.CreateAutonomousDatabaseRequest{
		CreateAutonomousDatabaseDetails: oci_database.CreateAutonomousDatabaseDetails{
			CompartmentId:        &compartmentId,
			CpuCoreCount:         &cpuCoreCount,
			DataStorageSizeInTBs: &dataStorageSizeInTbs,
			AdminPassword:        &adminPassword,
			DbName:               &dbName,
			DisplayName:          &displayName,
			IsDedicated:          &isDedicated,
			DbVersion:            &dbVersion,
		},
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: tfresource.GetRetryPolicy(false, "database"),
		},
	})
	if err != nil {
		return "", "", fmt.Errorf("[WARN] failed to Create adb with the error %v", err)
	}
	retryPolicy := tfresource.GetRetryPolicy(false, "database")
	retryPolicy.ShouldRetryOperation = acctest.ConditionShouldRetry(time.Duration(10*time.Minute), adbSweepWaitCondition, "database", false)

	if err != nil {
		return "", "", fmt.Errorf("[WARN] wait for adbSweepWaitCondition failed for %s resource with error %v", *createAdbResponse.Id, err)
	} else {
		log.Printf("[INFO] end of WaitTillCondition for resource %s ", *createAdbResponse.Id)
	}

	return *createAdbResponse.Id, *createAdbResponse.DbName, nil
}

func adbSweepWaitCondition(response oci_common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if autonomousDatabaseResponse, ok := response.Response.(oci_database.GetAutonomousDatabaseResponse); ok {
		return autonomousDatabaseResponse.LifecycleState == oci_database.AutonomousDatabaseLifecycleStateProvisioning
	}
	return false
}

func updateAdbInRegion(clients *tf_client.OracleClients, region string, autonomousDatabaseId string) (string, error) {
	databaseClient, err := createDatabaseClient(clients, region)
	if err != nil {
		return "", err
	}

	cpuCoreCount := 2
	dataStorageSizeInTBs := 2

	updateAdbResponse, err := databaseClient.UpdateAutonomousDatabase(context.Background(), oci_database.UpdateAutonomousDatabaseRequest{
		AutonomousDatabaseId: &autonomousDatabaseId,
		UpdateAutonomousDatabaseDetails: oci_database.UpdateAutonomousDatabaseDetails{
			CpuCoreCount:         &cpuCoreCount,
			DataStorageSizeInTBs: &dataStorageSizeInTBs,
		},
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: tfresource.GetRetryPolicy(false, "database"),
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] failed to update source autonomous database with the error %v", err)
	}
	return *updateAdbResponse.Id, nil
}

func changeDisasterRecoveryConfiguration(clients *tf_client.OracleClients, region string, autonomousDatabaseId string, disasterRecoveryType oci_database.ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum) (string, error) {
	databaseClient, err := createDatabaseClient(clients, region)
	if err != nil {
		return "", err
	}

	changeDrConfigResponse, err := databaseClient.ChangeDisasterRecoveryConfiguration(context.Background(), oci_database.ChangeDisasterRecoveryConfigurationRequest{
		AutonomousDatabaseId: &autonomousDatabaseId,
		ChangeDisasterRecoveryConfigurationDetails: oci_database.ChangeDisasterRecoveryConfigurationDetails{
			DisasterRecoveryType: disasterRecoveryType,
		},
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: tfresource.GetRetryPolicy(false, "database"),
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] failed to change disaster recovery type of standby autonomous database with the error %v", err)
	}
	return *changeDrConfigResponse.Id, nil
}

func changeSnapshotStandby(clients *tf_client.OracleClients, region string, autonomousDatabaseId string, isSnapshotStandby *bool) (string, error) {
	databaseClient, err := createDatabaseClient(clients, region)
	if err != nil {
		return "", err
	}

	changeDrConfigResponse, err := databaseClient.ChangeDisasterRecoveryConfiguration(context.Background(), oci_database.ChangeDisasterRecoveryConfigurationRequest{
		AutonomousDatabaseId: &autonomousDatabaseId,
		ChangeDisasterRecoveryConfigurationDetails: oci_database.ChangeDisasterRecoveryConfigurationDetails{
			IsSnapshotStandby: isSnapshotStandby,
		},
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: tfresource.GetRetryPolicy(false, "database"),
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] failed to connect/disconnect to snapshot standby with the error %v", err)
	}
	return *changeDrConfigResponse.Id, nil
}

func replicateBackupsStandby(clients *tf_client.OracleClients, region string, autonomousDatabaseId string, isReplicateBackupsEnabled *bool) (string, error) {
	databaseClient, err := createDatabaseClient(clients, region)
	if err != nil {
		return "", err
	}

	changeDrConfigResponse, err := databaseClient.ChangeDisasterRecoveryConfiguration(context.Background(), oci_database.ChangeDisasterRecoveryConfigurationRequest{
		AutonomousDatabaseId: &autonomousDatabaseId,
		ChangeDisasterRecoveryConfigurationDetails: oci_database.ChangeDisasterRecoveryConfigurationDetails{
			IsReplicateAutomaticBackups: isReplicateBackupsEnabled,
		},
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: tfresource.GetRetryPolicy(false, "database"),
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] failed to enable/disable replicate auto backups for standby with the error %v", err)
	}
	return *changeDrConfigResponse.Id, nil
}

func deleteAdbInRegion(clients *tf_client.OracleClients, region string, autonomousDatabaseId string) error {
	databaseClient, err := createDatabaseClient(clients, region)
	if err != nil {
		return err
	}

	if autonomousDatabaseId != "" {
		deleteAutonomousRequest := oci_database.DeleteAutonomousDatabaseRequest{}
		deleteAutonomousRequest.AutonomousDatabaseId = &autonomousDatabaseId

		_, err := databaseClient.DeleteAutonomousDatabase(context.Background(), deleteAutonomousRequest)
		if err != nil {
			return fmt.Errorf("failed to delete source autonomous database resource with error : %v", err)
		}
	}
	return nil
}

func adbWaitTillLifecycleStateAvailableCondition(response oci_common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if autonomousDatabaseResponse, ok := response.Response.(oci_database.GetAutonomousDatabaseResponse); ok {
		return autonomousDatabaseResponse.LifecycleState != oci_database.AutonomousDatabaseLifecycleStateAvailable
	}
	return false
}

func adbWaitTillLifecycleStateStandbyCondition(response oci_common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if autonomousDatabaseResponse, ok := response.Response.(oci_database.GetAutonomousDatabaseResponse); ok {
		return autonomousDatabaseResponse.LifecycleState != oci_database.AutonomousDatabaseLifecycleStateAvailable
	}
	return false
}

func getAdbFromSourceRegion(client *tf_client.OracleClients, resourceId *string, retryPolicy *oci_common.RetryPolicy) error {
	databaseClient, err := createDatabaseClient(client, utils.GetEnvSettingWithBlankDefault("source_region"))

	_, err = databaseClient.GetAutonomousDatabase(context.Background(), oci_database.GetAutonomousDatabaseRequest{
		AutonomousDatabaseId: resourceId,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func getAdbFromCurrentRegion(client *tf_client.OracleClients, resourceId *string, retryPolicy *oci_common.RetryPolicy) error {
	databaseClient, err := createDatabaseClient(client, utils.GetEnvSettingWithBlankDefault("region"))

	_, err = databaseClient.GetAutonomousDatabase(context.Background(), oci_database.GetAutonomousDatabaseRequest{
		AutonomousDatabaseId: resourceId,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}

func triggerSwitchoverOnAdbInRegion(clients *tf_client.OracleClients, region string, autonomousDatabaseId string, peerAdbId string) error {
	databaseClient, err := createDatabaseClient(clients, region)
	if err != nil {
		return err
	}

	if autonomousDatabaseId != "" {
		switchoverAutonomousRequest := oci_database.SwitchoverAutonomousDatabaseRequest{}
		switchoverAutonomousRequest.AutonomousDatabaseId = &autonomousDatabaseId
		switchoverAutonomousRequest.PeerDbId = &peerAdbId

		_, err := databaseClient.SwitchoverAutonomousDatabase(context.Background(), switchoverAutonomousRequest)
		if err != nil {
			return fmt.Errorf("failed to delete source autonomous database resource with error : %v", err)
		}
	}

	return nil
}

func createDatabaseClient(clients *tf_client.OracleClients, region string) (client oci_database.DatabaseClient, err error) {
	databaseClient, err := oci_database.NewDatabaseClientWithConfigurationProvider(*clients.DatabaseClient().ConfigurationProvider())
	if err != nil {
		return client, fmt.Errorf("cannot create client for the source region %s: %v", region, err)
	}
	err = tf_client.ConfigureClientVar(&databaseClient.BaseClient)
	if err != nil {
		return client, fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	databaseClient.SetRegion(region)
	return databaseClient, nil
}
