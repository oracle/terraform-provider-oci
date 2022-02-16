// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_database "github.com/oracle/oci-go-sdk/v58/database"
	oci_identity "github.com/oracle/oci-go-sdk/v58/identity"
)

func createAdbInRegion(clients *tf_client.OracleClients, region string) (string, error) {
	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")

	databaseClient, err := oci_database.NewDatabaseClientWithConfigurationProvider(*clients.DatabaseClient().ConfigurationProvider())
	if err != nil {
		return "", fmt.Errorf("cannot Create client for the source region %s: %v", region, err)
	}
	err = tf_client.ConfigureClientVar(&databaseClient.BaseClient)
	if err != nil {
		return "", fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	databaseClient.SetRegion(region)
	identityClient, err := oci_identity.NewIdentityClientWithConfigurationProvider(*clients.IdentityClient().ConfigurationProvider())
	if err != nil {
		return "", fmt.Errorf("cannot Create client for the source region %s: %v", region, err)
	}
	err = tf_client.ConfigureClientVar(&identityClient.BaseClient)
	if err != nil {
		return "", fmt.Errorf("cannot configure client for the source region: %v", err)
	}
	identityClient.SetRegion(region)

	cpu_core_count := 1
	data_storage_size_in_tbs := 1
	admin_password := string("BEstrO0ng_#11")
	db_name := utils.RandomString(13, utils.Charset)
	display_name := string("example_cr_source")
	is_dedicated := false

	createAdbResponse, err := databaseClient.CreateAutonomousDatabase(context.Background(), oci_database.CreateAutonomousDatabaseRequest{
		CreateAutonomousDatabaseDetails: oci_database.CreateAutonomousDatabaseDetails{
			CompartmentId:        &compartmentId,
			CpuCoreCount:         &cpu_core_count,
			DataStorageSizeInTBs: &data_storage_size_in_tbs,
			AdminPassword:        &admin_password,
			DbName:               &db_name,
			DisplayName:          &display_name,
			IsDedicated:          &is_dedicated,
		},
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: tfresource.GetRetryPolicy(false, "database"),
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] failed to Create adb with the error %v", err)
	}
	retryPolicy := tfresource.GetRetryPolicy(false, "database")
	retryPolicy.ShouldRetryOperation = acctest.ConditionShouldRetry(time.Duration(10*time.Minute), adbSweepWaitCondition, "database", false)

	if err != nil {
		return "", fmt.Errorf("[WARN] wait for adbSweepWaitCondition failed for %s resource with error %v", *createAdbResponse.Id, err)
	} else {
		log.Printf("[INFO] end of WaitTillCondition for resource %s ", *createAdbResponse.Id)
	}

	return *createAdbResponse.Id, nil
}

func adbSweepWaitCondition(response oci_common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if autonomousDatabaseResponse, ok := response.Response.(oci_database.GetAutonomousDatabaseResponse); ok {
		return autonomousDatabaseResponse.LifecycleState == oci_database.AutonomousDatabaseLifecycleStateProvisioning
	}
	return false
}
