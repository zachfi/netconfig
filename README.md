# Netconfig

[![Build Status](https://travis-ci.com/xaque208/netconfig.svg?branch=master)](https://travis-ci.com/xaque208/netconfig)

A simple template loader for network devices.  Currently supports Junos devices
using LDAP for inventory.

## Configuration

Below is the base configuration for the project.  You can specify its location
during execution with `--config`, or it will default to `~/.netconfig.yaml` of
the current user.

```
ldap:
  binddn: "cn=ldapuser,ou=services,dc=example,dc=com"
  bindpw: "secret"
  basedn: "dc=example,dc=com"
  host: "ldap.example.com"

junos:
  username: "netconfig"
  keyfile: "/home/user/.ssh/id_ed25519"

netconfig:
  configdir: "/home/zach/Org/n3kl/network"
```

## Data configuration

Here we will describe the use of the `data.yaml`, which is located within the
`configdir` specified in the above configuration file.

### YAML Data Loading

```yaml
data_dir: 'data'
hierarchy:
  - 'global.yaml'
  - 'role/{{ .Role }}.yaml'
  - 'host/{{ .Name }}.yaml'
```

The entries of the `hierarchy` are themselves templates that are rendered using
the current device.  Data is loaded in order, with subsequent data being merged
and replacing conflicting keys.  This results in the later data taking
precedence.  If ordered from most general to most specific as above, the paths
have the affect of grouping the data where it makes the most sense, allowing
de-duplication of data by using data common to devices at the correct tier.

### Template rendering

The following section in the `data.yaml` handles where to look for the
templates for a given device.  The `templtaes_dir` specifies the name of the
template directory with respect to the `configdir` of the specified
configuration file.  Each of the paths in `template_paths` are themselves
templates, that are rendered with the current device object.  All files ending
in `tmpl` found within these paths are rendered and loaded to the device.

```
template_dir: 'templates'
template_paths:
  - 'platform/{{ .Platform }}'
  - 'role/{{ .Role }}'
```
