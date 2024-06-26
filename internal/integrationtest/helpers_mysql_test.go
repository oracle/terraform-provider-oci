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
	oci_core "github.com/oracle/oci-go-sdk/v65/core"
	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"
	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"
)

func createDbSystemInRegion(clients *tf_client.OracleClients, region string) (string, error) {
	compartment := utils.GetEnvSettingWithBlankDefault("compartment_ocid")

	dbSystemClient, err := oci_mysql.NewDbSystemClientWithConfigurationProvider(*clients.DbSystemClient().ConfigurationProvider())
	if err != nil {
		return "", fmt.Errorf("cannot Create client for the region %s: %v", region, err)
	}

	err = tf_client.ConfigureClientVar(&dbSystemClient.BaseClient)
	if err != nil {
		return "", fmt.Errorf("cannot configure client for the region: %v", err)
	}

	dbSystemClient.SetRegion(region)

	identityClient, err := oci_identity.NewIdentityClientWithConfigurationProvider(*clients.IdentityClient().ConfigurationProvider())
	if err != nil {
		return "", fmt.Errorf("cannot Create client for the region %s: %v", region, err)
	}

	err = tf_client.ConfigureClientVar(&identityClient.BaseClient)
	if err != nil {
		return "", fmt.Errorf("cannot configure client for the region: %v", err)
	}

	identityClient.SetRegion(region)

	listAvailabilityDomainsResponse, err := identityClient.ListAvailabilityDomains(context.Background(),
		oci_identity.ListAvailabilityDomainsRequest{
			CompartmentId: &compartment,
		})
	if err != nil {
		return "", fmt.Errorf("[WARN] failed to get the domain name with the error %v", err)
	}
	domain := listAvailabilityDomainsResponse.Items[0].Name

	// Create subnet
	networkClient, err := oci_core.NewVirtualNetworkClientWithConfigurationProvider(*clients.VirtualNetworkClient().ConfigurationProvider())
	if err != nil {
		return "", fmt.Errorf("[WARN] cannot configure client for the region %v", err)
	}

	cidrBlockVcn := cidrBlockVcn
	networkClient.SetRegion(region)
	createVcnResponse, err := networkClient.CreateVcn(context.Background(), oci_core.CreateVcnRequest{
		CreateVcnDetails: oci_core.CreateVcnDetails{
			CidrBlock:     &cidrBlockVcn,
			CompartmentId: &compartment,
		}})

	if err != nil {
		return "", fmt.Errorf("[WARN] failed to Create source VCN with the error %v", err)
	}

	cidrBlockSubnet := cidrBlockSubnet

	createSubnetResponse, err := networkClient.CreateSubnet(context.Background(), oci_core.CreateSubnetRequest{
		CreateSubnetDetails: oci_core.CreateSubnetDetails{
			CompartmentId: &compartment,
			CidrBlock:     &cidrBlockSubnet,
			VcnId:         createVcnResponse.Id,
		},
	})

	// Required DB System configurations
	shape := "MySQL.VM.Standard.E3.1.8GB"
	adminPassword := "BEstrO0ng_#11"
	adminUsername := "adminUser"

	createDbSystemResponse, err := dbSystemClient.CreateDbSystem(context.Background(), oci_mysql.CreateDbSystemRequest{
		CreateDbSystemDetails: oci_mysql.CreateDbSystemDetails{
			SubnetId:           createSubnetResponse.Id,
			ShapeName:          &shape,
			AvailabilityDomain: domain,
			CompartmentId:      &compartment,
			AdminUsername:      &adminUsername,
			AdminPassword:      &adminPassword,
		},
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: tfresource.GetRetryPolicy(false, "mysql"),
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] failed to Create source DB System with the error %v", err)
	}

	retryPolicy := tfresource.GetRetryPolicy(false, "mysql")
	retryPolicy.ShouldRetryOperation = acctest.ConditionShouldRetry(15*time.Minute, dbSystemActiveWaitCondition, "mysql", false)

	_, err = dbSystemClient.GetDbSystem(context.Background(), oci_mysql.GetDbSystemRequest{
		DbSystemId: createDbSystemResponse.Id,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] wait for dbSystemActiveWaitCondition failed for %s resource with error %v", *createDbSystemResponse.Id, err)
	} else {
		log.Printf("[INFO] end of WaitTillCondition for resource %s ", *createDbSystemResponse.Id)
	}

	return *createDbSystemResponse.Id, nil
}

func createBackupInRegion(clients *tf_client.OracleClients, region string, dbSystemId *string) (string, error) {
	dbBackupsClient, err := initializeDbBackupsClientWithConfigurationProvider(clients, region, "DbSystemClient")

	createMysqlBackupResponse, err := dbBackupsClient.CreateBackup(context.Background(), oci_mysql.CreateBackupRequest{
		CreateBackupDetails: oci_mysql.CreateBackupDetails{
			DbSystemId: dbSystemId,
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] failed to Create source DB System Backup with the error %v", err)

	}

	retryPolicy := tfresource.GetRetryPolicy(false, "core")
	retryPolicy.ShouldRetryOperation = acctest.ConditionShouldRetry(10*time.Minute, mysqlBackupActiveWaitCondition, "mysql", false)
	_, err = dbBackupsClient.GetBackup(context.Background(), oci_mysql.GetBackupRequest{
		BackupId: createMysqlBackupResponse.Id,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	if err != nil {
		return "", fmt.Errorf("[WARN] wait for mysqlBackupActiveWaitCondition failed for %s resource with error %v", *createMysqlBackupResponse.Id, err)
	} else {
		log.Printf("[INFO] end of WaitTillCondition for resource %s ", *createMysqlBackupResponse.Id)
	}

	return *createMysqlBackupResponse.Id, nil
}

func deleteBackupInRegion(clients *tf_client.OracleClients, region string, mysqlBackupId string) error {
	dbBackupsClient, _ := initializeDbBackupsClientWithConfigurationProvider(clients, region, "DbBackupsClient")

	if mysqlBackupId != "" {
		deleteMysqlBackupRequest := oci_mysql.DeleteBackupRequest{}
		deleteMysqlBackupRequest.BackupId = &mysqlBackupId

		_, err := dbBackupsClient.DeleteBackup(context.Background(), deleteMysqlBackupRequest)
		if err != nil {
			return fmt.Errorf("failed to delete source DB System Backup %s resource with error %v", *deleteMysqlBackupRequest.BackupId, err)
		}
	}

	return nil
}

func deleteDbSystemInRegion(clients *tf_client.OracleClients, region string, mysqlDbSystemId string) error {
	dbSystemClient, err := oci_mysql.NewDbSystemClientWithConfigurationProvider(*clients.DbSystemClient().ConfigurationProvider())
	if err != nil {
		return fmt.Errorf("cannot Create client for the region %s: %v", region, err)
	}

	err = tf_client.ConfigureClientVar(&dbSystemClient.BaseClient)
	if err != nil {
		return fmt.Errorf("cannot configure client for the region: %v", err)
	}

	dbSystemClient.SetRegion(region)

	if mysqlDbSystemId != "" {
		deleteMysqlDbSystemRequest := oci_mysql.DeleteDbSystemRequest{}
		deleteMysqlDbSystemRequest.DbSystemId = &mysqlDbSystemId

		_, err := dbSystemClient.DeleteDbSystem(context.Background(), deleteMysqlDbSystemRequest)
		if err != nil {
			return fmt.Errorf("failed to delete source DB System %s resource with error %v", *deleteMysqlDbSystemRequest.DbSystemId, err)
		}
	}

	return nil
}

func dbSystemActiveWaitCondition(response oci_common.OCIOperationResponse) bool {
	if dbSystemResponse, ok := response.Response.(oci_mysql.GetDbSystemResponse); ok {
		return dbSystemResponse.LifecycleState != oci_mysql.DbSystemLifecycleStateActive
	}

	return false
}

func mysqlBackupActiveWaitCondition(response oci_common.OCIOperationResponse) bool {
	if dbBackupResponse, ok := response.Response.(oci_mysql.GetBackupResponse); ok {
		return dbBackupResponse.LifecycleState != oci_mysql.BackupLifecycleStateActive
	}

	return false
}

func initializeDbBackupsClientWithConfigurationProvider(clients *tf_client.OracleClients, region string, clientType string) (*oci_mysql.DbBackupsClient, error) {
	var dbBackupsClient oci_mysql.DbBackupsClient
	var err error

	switch clientType {
	case "DbSystemClient":
		dbBackupsClient, err = oci_mysql.NewDbBackupsClientWithConfigurationProvider(*clients.DbSystemClient().ConfigurationProvider())
	case "DbBackupsClient":
		dbBackupsClient, err = oci_mysql.NewDbBackupsClientWithConfigurationProvider(*clients.DbBackupsClient().ConfigurationProvider())
	default:
		return nil, fmt.Errorf("unsupported client type: %s", clientType)
	}

	if err != nil {
		return nil, fmt.Errorf("cannot create client for the region %s: %v", region, err)
	}

	err = tf_client.ConfigureClientVar(&dbBackupsClient.BaseClient)
	if err != nil {
		return nil, fmt.Errorf("cannot configure client for the region: %v", err)
	}

	dbBackupsClient.SetRegion(region)

	return &dbBackupsClient, err
}
