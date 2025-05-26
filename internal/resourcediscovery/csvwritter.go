package resourcediscovery

import (
	"encoding/csv"
	"os"
	"path/filepath"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	tf_export "github.com/oracle/terraform-provider-oci/internal/commonexport"
)

// New function to write resources to CSV
func writeResourcesToCSV(resources []*tf_export.OCIResource, compartmentID, outputDir string) error {
	// Create CSV file
	utils.Logf("Writing %d resources to CSV", len(resources))
	file, err := os.Create(filepath.Join(outputDir, "exported_resources.csv"))
	if err != nil {
		return err
	}
	defer file.Close()

	// Initialize CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV headers
	headers := []string{"ResourceOCID", "ResourceType", "CompartmentID", "Name", "TerraformReference", "State", "CreatedBy", "TimeCreated", "TimeUpdated"}
	if err := writer.Write(headers); err != nil {
		return err
	}

	// Write resource data
	for _, resource := range resources {
		// Extract common attributes; add more as needed
		name := resource.TerraformName
		state := "unknown"
		createdAt := ""
		updatedAt := ""
		createdBy := ""
		if timeCreatedAttr, ok := resource.SourceAttributes["time_created"]; ok {
			createdAt = timeCreatedAttr.(string)
		}
		if stateAttr, ok := resource.SourceAttributes["state"]; ok {
			state = stateAttr.(string)
		}
		if updatedAtattr, ok := resource.SourceAttributes["time_updated"]; ok {
			updatedAt = updatedAtattr.(string)
		}
		if definedTags, ok := resource.SourceAttributes["defined_tags"]; ok {
			definedTagsMap := definedTags.(map[string]interface{})
			if oracleTags, ok := definedTagsMap["oracle_tags"]; ok {
				oracleTagsMap := oracleTags.(map[string]interface{})
				if created_by, ok := oracleTagsMap["CreatedBy"]; ok {
					createdBy = created_by.(string)
				}
			}
		}
		record := []string{
			resource.Id,
			resource.TerraformClass,
			compartmentID,
			name,
			resource.GetTerraformReference(),
			state,
			createdBy,
			createdAt,
			updatedAt,
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	utils.Logf("csv file %s generated for %d resources", filepath.Join(outputDir, "exported_resources.csv"), len(resources))
	return nil
}
