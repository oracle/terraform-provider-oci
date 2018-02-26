// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package oci_tool

import (
	"fmt"
	"os"
)

// Copy target directory and append .backup
func CreateBackup(targetDir string, backupDir string) (err error) {
	fmt.Println("Creating backup...", targetDir, "-->", backupDir)

	fi, err := os.Stat(targetDir)

	if err != nil {
		return fmt.Errorf("Error reading directory\n %s", err)
	}

	if !fi.IsDir() {
		return fmt.Errorf("File targeted for migration")
	}

	_, err = os.Stat(backupDir)

	if err == nil {
		return fmt.Errorf("Attempting to overwrite backups")
	}

	fmt.Println("Copying", targetDir, "-->", backupDir)

	err = ProcessDirectory(targetDir, backupDir, CopyFile)

	if err != nil {
		return err
	}

	bfi, err := os.Stat(backupDir)

	if fi.Size() != bfi.Size() {
		return fmt.Errorf("Backup corrupt")
	}

	fmt.Println("Complete")
	return
}

// Overwrite target directory with contents of .backup directory
func RestoreBackup(backupDir string, targetDir string) (err error) {
	fmt.Println("Restoring from backup...")

	fi, err := os.Stat(backupDir)

	if err != nil {
		return fmt.Errorf("Error reading backup\n %s", err)
	}

	err = os.RemoveAll(targetDir)

	if err != nil {
		return fmt.Errorf("Error removing original directory\n %s", err)
	}

	os.MkdirAll(targetDir, fi.Mode())

	err = ProcessDirectory(backupDir, targetDir, CopyFile)

	if err != nil {
		return fmt.Errorf("Error restoring from backup directory\n %s", err)
	}

	fmt.Println("Complete")
	return
}

// Remove .backup directory
func DeleteBackup(backupDir string) (err error) {
	fmt.Println("Purging backup...")

	err = os.RemoveAll(backupDir)

	if err != nil {
		return fmt.Errorf("Error removing backup directory\n %s", err)
	}

	fmt.Println("Complete")
	return
}

// Traverse all .tf files and apply transforms
func Migrate(targetDir string, backupDir string) (err error) {
	fmt.Println("Migrating plan directory...")
	err = CreateBackup(targetDir, backupDir)

	if err != nil {
		return fmt.Errorf("Error backing up directory before migration\n %s", err)
	}

	err = ProcessDirectory(targetDir, backupDir, MigratePlanFile, ".tf", ".tfstate")

	if err != nil {
		return fmt.Errorf("Error removing backup directory\n %s", err)
	}

	fmt.Println("Complete")
	return
}

// Traverse all .tf files and insert `region = "us-phoenix-1"` where region is not specified
func AddRegionField(targetDir string, backupDir string) (err error) {
	fmt.Println("Scanning plans for providers missing region value...")

	err = ProcessDirectory(targetDir, backupDir, AddRegionToProvider, ".tf")

	if err != nil {
		return fmt.Errorf("Error scanning providers for missing region value\n %s", err)
	}

	fmt.Println("Complete")
	return
}
