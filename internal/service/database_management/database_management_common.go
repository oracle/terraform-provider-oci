package database_management

import (
	"fmt"

	oci_database_management "github.com/oracle/oci-go-sdk/v65/databasemanagement"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

// Choose feature if someone has different resource block with
// enable_database_dbm_feature = false and feaure=oci_database_management.DbManagementFeatureEnum
func resolveFeatureForDBOperation(s *tfresource.BaseCrud) oci_database_management.DbManagementFeatureEnum {
	var chosen_feature oci_database_management.DbManagementFeatureEnum
	if feature, ok := s.D.GetOkExists("feature"); ok {
		tmp := feature.(string)
		chosen_feature = oci_database_management.DbManagementFeatureEnum(tmp)
	}

	if len(chosen_feature) == 0 {
		// Choose feature if someone calls terrafom destroy
		if featureDetails, ok := s.D.GetOkExists("feature_details"); ok {
			if tmpList := featureDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "feature_details", 0)
				featureRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "feature"))
				if ok {
					chosen_feature = oci_database_management.DbManagementFeatureEnum(featureRaw.(string))
				} else {
					chosen_feature = "" // default value
				}
			}
		}
	}
	return chosen_feature
}
