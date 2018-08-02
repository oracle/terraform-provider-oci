#!/usr/bin/env python

#
# This script manipulates pom.xml to either add new specs or update the versions of existing specs.
#

import xml.etree.ElementTree as ET
import re
import click
from click.exceptions import UsageError

DEFAULT_POM_LOCATION = "../pom.xml"
DEFAULT_GITHUB_WHITELIST_LOCATION = "../github.whitelist"
DEFAULT_MAKE_FILE_LOCATION = "../Makefile"

PROPERTIES_ELEMENT_ARTIFACT_VERSION = """<{artifact_id}.artifact.version>{version}</{artifact_id}.artifact.version>"""
PROPERTIES_ELEMENT_ARTIFACT_ID = """<{artifact_id}.artifact.id>{artifact_id}</{artifact_id}.artifact.id>"""
PROPERTIES_ELEMENT_SPEC_NAME = """<{artifact_id}.spec.name>{spec_path_relative_to_jar}</{artifact_id}.spec.name>"""

UNPACK_EXECUTION_TEMPLATE = """
<execution>
    <id>unpack-{spec_name}</id>
    <phase>initialize</phase>
    <goals>
        <goal>unpack</goal>
    </goals>
    <configuration>
        <artifactItems>
            <artifactItem>
                <groupId>{group_id}</groupId>
                <artifactId>${{{artifact_id}.artifact.id}}</artifactId>
                <type>jar</type>
                <includes>**/*</includes>
                <outputDirectory>${{spec.temp.dir}}/{spec_name}</outputDirectory>
            </artifactItem>
        </artifactItems>
    </configuration>
</execution>
"""

PREFER_EXECUTION_TEMPLATE = """
<execution>
    <id>spec-conditionals-prefer-{spec_name}</id>
    <phase>initialize</phase>
    <goals>
        <goal>prefer</goal>
    </goals>
    <configuration>
        <inputFiles>
            <!-- New layout: source/<spec.proto.yaml> -->
            <inputFile>${{spec.temp.dir}}/{spec_name}/source/${{{artifact_id}.spec.name}}</inputFile>
            <!-- Old layout: ./<spec.proto.yaml> -->
            <inputFile>${{spec.temp.dir}}/{spec_name}/${{{artifact_id}.spec.name}}</inputFile>
        </inputFiles>
        <outputFile>${{preferred.temp.dir}}/${{{artifact_id}.spec.name}}</outputFile>
    </configuration>
</execution>
"""

PREPROCESS_EXECUTION_TEMPLATE = """
<execution>
    <id>spec-conditionals-preprocess-{spec_name}</id>
    <phase>initialize</phase>
    <goals>
        <goal>preprocess</goal>
    </goals>
    <configuration>
        <inputFile>${{preferred.temp.dir}}/${{{artifact_id}.spec.name}}</inputFile>
        <outputFile>${{preprocessed.temp.dir}}/${{{artifact_id}.spec.name}}</outputFile>
        <groupFile>${{enabled.groups.file}}</groupFile>
    </configuration>
</execution>
"""

GENERATE_EXECUTION_TEMPLATE = """
<execution>
    <id>go-public-sdk-{spec_name}</id>
    <phase>compile</phase>
    <goals>
        <goal>generate</goal>
    </goals>
    <configuration>
        <language>oracle-go-sdk</language>
        <specPath>${{preprocessed.temp.dir}}/${{{artifact_id}.spec.name}}</specPath>
        <outputDir>${{env.GOPATH}}/src/${{fullyQualifiedProjectName}}</outputDir>
        <basePackage>{spec_name}</basePackage>
        <specGenerationType>${{generationType}}</specGenerationType>
        <additionalProperties>
            <specName>{spec_name}</specName>
            <fqProjectName>${{fullyQualifiedProjectName}}</fqProjectName>
            <serviceHostName>{subdomain}</serviceHostName>
            {regional_non_regional_service_overrides}
        </additionalProperties>
        <featureIdConfigFile>${{project.basedir}}/featureId.yaml</featureIdConfigFile>
    </configuration>
</execution>
"""

CLEAN_ELEMENT_TEMPLATE = """
 <fileset>
    <directory>lib/oci/{spec_name}</directory>
    <includes>
        <include>**/*</include>
    </includes>
   <excludes>
	<exclude>util.rb</exclude>
   </excludes>
</fileset>
"""

DEPENDENCY_MANAGEMENT_TEMPLATE = """
    <dependency>
        <groupId>{group_id}</groupId>
        <artifactId>${{{artifact_id}.artifact.id}}</artifactId>
        <version>${{{artifact_id}.artifact.version}}</version>
    </dependency>
"""

ns = {"ns":"http://maven.apache.org/POM/4.0.0"}

# allow default namespace for output, dont print ns0: prefixes everywhere
ET.register_namespace('',"http://maven.apache.org/POM/4.0.0")

def parse_pom(pom_location):
    return ET.parse(pom_location)


def generate_and_add_property_element(pom, artifact_id, version, spec_path_relative_to_jar):
    artifact_version_content = PROPERTIES_ELEMENT_ARTIFACT_VERSION.format(
        artifact_id=artifact_id,
        version=version
    )
    artifact_version_element = ET.fromstring(artifact_version_content)

    artifact_id_content = PROPERTIES_ELEMENT_ARTIFACT_ID.format(
        artifact_id=artifact_id
    )
    artifact_id_element = ET.fromstring(artifact_id_content)

    spec_name_content = PROPERTIES_ELEMENT_SPEC_NAME.format(
        artifact_id=artifact_id,
        spec_path_relative_to_jar=spec_path_relative_to_jar
    )
    spec_name_element = ET.fromstring(spec_name_content)

    xpath = ".//ns:properties"
    properties = pom.findall(xpath, ns)[0]
    properties.append(artifact_version_element)
    properties.append(artifact_id_element)
    properties.append(spec_name_element)


def update_relative_spec_path(pom, artifact_id, spec_path_relative_to_jar):
    xpath = ".//ns:properties/ns:{artifact_id}.spec.name".format(artifact_id=artifact_id)
    spec_file_node = pom.findall(xpath, ns)[0]
    spec_file_node.text = spec_path_relative_to_jar


def generate_and_add_unpack_element(pom, spec_name, group_id, artifact_id, spec_path_relative_to_jar):
    content = UNPACK_EXECUTION_TEMPLATE.format(
        spec_name=spec_name,
        group_id=group_id,
        artifact_id=artifact_id,
        spec_path_relative_to_jar=spec_path_relative_to_jar)

    unpack_element = ET.fromstring(content)

    # find maven-dependency-plugin where unpacking happens
    unpack_plugin_executions = pom.findall(".//ns:plugin[ns:artifactId='maven-dependency-plugin']/ns:executions", ns)[0]
    unpack_plugin_executions.append(unpack_element)


def generate_and_add_prefer_element(pom, spec_name, group_id, artifact_id, spec_path_relative_to_jar):
    content = PREFER_EXECUTION_TEMPLATE.format(
        spec_name=spec_name,
        group_id=group_id,
        artifact_id=artifact_id,
        spec_path_relative_to_jar=spec_path_relative_to_jar)

    unpack_element = ET.fromstring(content)

    # find maven-dependency-plugin where unpacking happens
    unpack_plugin_executions = pom.findall(".//ns:plugin[ns:artifactId='spec-conditionals-preprocessor-plugin']/ns:executions", ns)[0]
    unpack_plugin_executions.append(unpack_element)


def generate_and_add_preprocess_element(pom, spec_name, group_id, artifact_id, spec_path_relative_to_jar):
    content = PREPROCESS_EXECUTION_TEMPLATE.format(
        spec_name=spec_name,
        group_id=group_id,
        artifact_id=artifact_id,
        spec_path_relative_to_jar=spec_path_relative_to_jar)

    unpack_element = ET.fromstring(content)

    # find maven-dependency-plugin where unpacking happens
    unpack_plugin_executions = pom.findall(".//ns:plugin[ns:artifactId='spec-conditionals-preprocessor-plugin']/ns:executions", ns)[0]
    unpack_plugin_executions.append(unpack_element)


def generate_and_add_generate_section(pom, spec_name, artifact_id, spec_path_relative_to_jar, subdomain, regional_sub_service_overrides, non_regional_sub_service_overrides):
    regional_non_regional_service_overrides_content = ''
    if regional_sub_service_overrides or non_regional_sub_service_overrides:
        if regional_sub_service_overrides:
            for override in regional_sub_service_overrides:
                regional_non_regional_service_overrides_content += '<isRegionalClient.{service_name}>true</isRegionalClient.{service_name}>\n'.format(service_name=override)

        if non_regional_sub_service_overrides:
            for override in non_regional_sub_service_overrides:
                regional_non_regional_service_overrides_content += '<isRegionalClient.{service_name}>false</isRegionalClient.{service_name}>\n'.format(service_name=override)

    content = GENERATE_EXECUTION_TEMPLATE.format(
        artifact_id=artifact_id,
        spec_name=spec_name,
        spec_path_relative_to_jar=spec_path_relative_to_jar,
        subdomain=subdomain,
        regional_non_regional_service_overrides=regional_non_regional_service_overrides_content)

    generate_element = ET.fromstring(content)

    # find bmc-sdk-swagger-maven-plugin where generation happens
    generate_plugin_executions = pom.findall(".//ns:plugin[ns:artifactId='bmc-sdk-swagger-maven-plugin']/ns:executions", ns)[0]
    generate_plugin_executions.append(generate_element)


def generate_and_add_clean_section(pom, spec_name):
    content = CLEAN_ELEMENT_TEMPLATE.format(
        spec_name=spec_name)

    clean_element = ET.fromstring(content)

    # find filesetes where clean directory goes
    filesets = pom.findall(".//ns:plugin[ns:artifactId='maven-clean-plugin']//ns:filesets", ns)[0]
    filesets.append(clean_element)


def generate_and_add_dependency_management_section(pom, spec_name, group_id, artifact_id, version):
    content = DEPENDENCY_MANAGEMENT_TEMPLATE.format(
        spec_name=spec_name,
        group_id=group_id,
        artifact_id=artifact_id,
        version=version)

    dep_mgt_element = ET.fromstring(content)

    # find dependencies where version is specified
    dependencies = pom.findall(".//ns:dependencyManagement/ns:dependencies", ns)[0]
    dependencies.append(dep_mgt_element)


def update_version_of_existing_spec(pom, artifact_id, version):
    xpath = ".//ns:properties//ns:{artifact_id}.artifact.version".format(artifact_id=artifact_id)
    dependency = pom.findall(xpath, ns)[0]
    dependency.text = version


def indent(elem, level=0):
    indent_str = "    "
    i = "\n" + level*indent_str
    if len(elem):
        if not elem.text or not elem.text.strip():
            elem.text = i + indent_str
        for e in elem:
            indent(e, level+1)
            if not e.tail or not e.tail.strip():
                e.tail = i + indent_str
        if not e.tail or not e.tail.strip():
            e.tail = i
    else:
        if level and (not elem.tail or not elem.tail.strip()):
            elem.tail = i


def add_spec_module_to_github_whitelist(spec_name, github_whitelist_location):
    with open(github_whitelist_location, 'a') as f:
        f.write('\n^{}/'.format(spec_name))

def add_spec_name_to_make_file(spec_name, make_file_location):
    specNameToken = '##SPECNAME##'
    with open(make_file_location) as f:
        newText=f.read().replace(specNameToken, "{} {}".format(spec_name, specNameToken))

    with open(make_file_location, "w") as f:
        f.write(newText)

def goify_specname(name):
    return name.replace('_', '').lower()

def isNewService(pom, artifact_id):
    propertyXPath = ".//ns:properties"
    properties = pom.findall(propertyXPath, ns)[0]

    for child in properties.getiterator():
        if (child.text.lower() == artifact_id.lower()):
            return True
    return False

def add_or_update_spec(artifact_id=None, group_id=None, spec_name=None, relative_spec_path=None, subdomain=None, version=None, spec_generation_type=None, regional_sub_service_overrides=None, non_regional_sub_service_overrides=None, pom_location=None, github_whitelist_location=None):
    if not artifact_id:
        raise click.exceptions.MissingParameter(param_type='option', param_hint='--artifact-id')

    if spec_generation_type:
        print('Note: --spec-generation-type is ignored for the GO SDK, since it is set in the ../pom.xml file for all modules')

    pom = parse_pom(pom_location)

    # force format of spec_name by removing underscore and lower case
    if spec_name:
        spec_name = goify_specname(spec_name)

    # determine if this artifact is already in the spec
    if isNewService(pom, artifact_id):
        print('Artifact {} already exists in pom.xml. Updating specified fields...'.format(artifact_id))

        if version:
            update_version_of_existing_spec(pom, artifact_id, version)

        if relative_spec_path:
            update_relative_spec_path(pom, artifact_id, relative_spec_path)
    else:
        if not version:
            raise UsageError('Must specify --version for a new spec')

        if not subdomain:
            raise UsageError('Must specify --subdomain for a new spec')

        if not group_id:
            raise UsageError('Must specify --group-id for new spec')

        if not spec_name:
            raise UsageError('Must specify --spec-name for new spec')

        if not relative_spec_path:
            raise UsageError('Must specify --relative-spec-path for new spec')

        print('Artifact {} does not exist in pom.xml. Adding it...'.format(spec_name))
        generate_and_add_property_element(pom, artifact_id, version, relative_spec_path)
        generate_and_add_unpack_element(pom, spec_name, group_id, artifact_id, relative_spec_path)
        generate_and_add_prefer_element(pom, spec_name, group_id, artifact_id, relative_spec_path)
        generate_and_add_preprocess_element(pom, spec_name, group_id, artifact_id, relative_spec_path)
        generate_and_add_generate_section(pom, spec_name, artifact_id, relative_spec_path, subdomain, regional_sub_service_overrides, non_regional_sub_service_overrides)
        generate_and_add_clean_section(pom, spec_name)
        generate_and_add_dependency_management_section(pom, spec_name, group_id, artifact_id, version)
        add_spec_module_to_github_whitelist(spec_name, github_whitelist_location)
        add_spec_name_to_make_file(spec_name, DEFAULT_MAKE_FILE_LOCATION)

    # pretty print pom
    indent(pom.getroot())
    pom.write(pom_location, encoding="UTF-8", xml_declaration=True)

    print('Success!')


@click.command()
@click.option('--artifact-id', help='The artifact id for the spec artifact (e.g. coreservices-api-spec')
@click.option('--group-id', help='The group id for the spec artifact (e.g. com.oracle.pic.commons)')
@click.option('--spec-name', help='The name of the spec. This will be used (e.g. core, identity, object_storage). This is also used as the module name.')
@click.option('--relative-spec-path', help='The relative path of the spec within the artifact (e.g. coreservices-api-spec-20160918-external.yaml)')
@click.option('--subdomain', help='The subdomain for the service (e.g. if the endpoint is https://iaas.{domain}/20160918), the subdomain is "iaas"')
@click.option('--version', help='The version of the spec artifact (e.g. 0.0.1-SNAPSHOT')
@click.option('--spec-generation-type', help='The generation type: PUBLIC or PREVIEW')
@click.option('--regional-sub-service-overrides', multiple=True, help="""For specs that contain multiple services
(because there are operations with different tags in the spec), which of those services should be considered regional.
Services are considered as regional by default.

This should be the snake_cased name of the tag/service. For example kms_provisioning instead of kmsProvisioning.

This parameter can be provided multiple times""")
@click.option('--non-regional-sub-service-overrides', multiple=True, help="""For specs that contain multiple services
(because there are operations with different tags in the spec), which of those services should be considered non-regional.

This should be the snake_cased name of the tag/service. For example kms_provisioning instead of kmsProvisioning.

This parameter can be provided multiple times""")
@click.option('--pom-location', type=click.Path(exists=True), default=DEFAULT_POM_LOCATION, help='Location of the pom.xml file to update')
@click.option('--github-whitelist-location', type=click.Path(exists=True), default=DEFAULT_GITHUB_WHITELIST_LOCATION, help='Location of the github.whitelist file to update')
def add_or_update_spec_command(artifact_id, group_id, spec_name, relative_spec_path, subdomain, version, spec_generation_type, regional_sub_service_overrides, non_regional_sub_service_overrides, pom_location, github_whitelist_location):
    add_or_update_spec(artifact_id, group_id, spec_name, relative_spec_path, subdomain, version, spec_generation_type, regional_sub_service_overrides, non_regional_sub_service_overrides, pom_location, github_whitelist_location)


if __name__ == '__main__':
    add_or_update_spec_command()
