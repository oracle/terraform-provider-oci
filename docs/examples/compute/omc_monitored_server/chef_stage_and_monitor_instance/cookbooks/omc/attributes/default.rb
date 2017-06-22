default['omc']['os_user'] = "omc"
default['omc']['installer_group'] = "omcinstall"
default['omc']['os_timezone'] = "America/Lima"
default['omc']['install'] = "/omc/install"
default['omc']['stage'] = "/omc/stage"
default['omc']['app'] = "/omc/app"
default['omc']['apm'] = "/omc/apm"
default['omc']['regkey'] = "RvnWiqoF63rVq2ZW9G09Gu4W0N"
default['omc']['port'] = "1830"

default['omc']['packages'] = [
    { "name" => "binutils", "arch" => "x86_64" },
    { "name" => "perl","arch" => "x86_64" },
    { "name" => "unzip","arch" => "x86_64" },
    { "name" => "bind-utils","arch" => "x86_64" },
    { "name" => "bc","arch" => "x86_64" },
    { "name" => "rng-tools","arch" => "x86_64" }
]

default['authorization']['sudo']['include_sudoers_d'] = true

