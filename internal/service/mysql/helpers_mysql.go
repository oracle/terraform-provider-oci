package mysql

import (
	"fmt"

	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
)

func (s *MysqlMysqlBackupResourceCrud) createDbBackupClientInRegion(region string) error {
	if s.DestRegionClient == nil {
		dbBackupClient, err := oci_mysql.NewDbBackupsClientWithConfigurationProvider(*s.Client.ConfigurationProvider())
		if err != nil {
			return fmt.Errorf("cannot Create client for the region: %v", err)
		}
		err = tf_client.ConfigureClientVar(&dbBackupClient.BaseClient)
		if err != nil {
			return fmt.Errorf("cannot configure client for the region: %v", err)
		}
		s.DestRegionClient = &dbBackupClient
	}
	s.DestRegionClient.SetRegion(region)

	return nil
}
