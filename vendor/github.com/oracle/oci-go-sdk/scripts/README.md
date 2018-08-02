# go-sdk SDK Utilities

This README describes the various scripts and utilities included in this folder.

## add_or_update_spec.py
This simplifies the process of updating a pom.xml file for a new service, or changing the spec version for an existing service (it is more useful in the case of the former, rather than the latter).

### Setup
You should run this script in a Python virtual environment (see [here](https://bitbucket.aka.lgl.grungy.us/projects/SDK/repos/python-sdk/browse/Internal-README.rst) for more information on doing the basic setup for a virtual environment if you don't already have one).

Once you are in a virtual environment, you can run `pip install -r add_or_update_spec-requirements.txt` to install the dependencies for the script.

### Execution
You can run the script with `python add_or_update_spec.py --help` to get an overview of the various options/parameters which can be fed to it.

If you want to test the script out, this folder has a sample pom.xml file in it. Here is an example of running the script to add the Container Engine service to the pom.xml file:

```
python add_or_update_spec.py --artifact-id clusters-api-spec \
--group-id com.oracle.pic.clusters \
--spec-name container_engine \
--relative-spec-path clusters-api-spec.yaml \
--subdomain containerengine \
--version 1.0.7 \
--pom-location test_pom.xml
```