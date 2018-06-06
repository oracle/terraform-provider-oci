# Change Log
All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/).

## 0.2 - 2018-06-05

### Added
- *** IMPORTANT *** Updated process to use Instance Principal authorization method for OCI-CLI.  This
  removes the requirement to upload the user private key into the ipxe server.  *HOWEVER*, it does REQUIRE 
  the creation of a Dynamic Group for the Compartment this is being executed in.  This is included in the
  prerequisites section below.  The new process WILL NOT WORK without an Instance Principal (also known
  as a Dynamic Group).
- Changed our blank image OCID to use a new global value.  This makes the process *more* portable across
  regions (but not completely...see next bullet).
- Updated the ipxe server image list to include the LHR region.  LHR is now fully supported on this process.
  
### Fixed
- Fixed issue with OCI-CLI and pip that was preventing the process from working.  The original 
  process attempted to install OCI-CLI as a global resource in the global python repo...this broke
  after some changes in pip.  So now we install pip locally to each user.
- Removed requirement for sharutils/uuencode/uudecode to be installed on both TF client and ipxe server.  
  Replaced with base64 for portability.
- Fixed to work on both Mac and Linux TF clients.

## 0.1 - 2018-01-12

### Added
- Initial release.