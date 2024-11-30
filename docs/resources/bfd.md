---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_bfd Resource - terraform-provider-iosxr"
subcategory: "System"
description: |-
  This resource can manage the BFD configuration.
---

# iosxr_bfd (Resource)

This resource can manage the BFD configuration.

## Example Usage

```terraform
resource "iosxr_bfd" "example" {
  echo_disable                                           = true
  echo_latency_detect_percentage                         = 200
  echo_latency_detect_count                              = 10
  echo_startup_validate_force                            = true
  echo_ipv4_source                                       = "10.1.1.1"
  echo_ipv4_bundle_per_member_preferred_minimum_interval = 200
  trap_singlehop_pre_mapped                              = true
  multipath_locations = [
    {
      location_name = "0/0/CPU0"
    }
  ]
  multihop_ttl_drop_threshold            = 200
  dampening_initial_wait                 = 3600
  dampening_secondary_wait               = 3200
  dampening_maximum_wait                 = 3100
  dampening_threshold                    = 60000
  dampening_extensions_down_monitoring   = true
  dampening_disable                      = true
  dampening_bundle_member_l3_only_mode   = true
  dampening_bundle_member_initial_wait   = 5184
  dampening_bundle_member_secondary_wait = 6184
  dampening_bundle_member_maximum_wait   = 7184
  bundle_coexistence_bob_blb_inherit     = false
  bundle_coexistence_bob_blb_logical     = true
  interfaces = [
    {
      interface_name        = "GigabitEthernet0/0/0/0"
      echo_disable          = true
      echo_ipv4_source      = "12.1.1.1"
      ipv6_checksum_disable = true
      disable               = true
      local_address         = "33.33.31.1"
      tx_interval           = 3200
      rx_interval           = 4200
      multiplier            = 40
    }
  ]
  ipv6_checksum_disable = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `bundle_coexistence_bob_blb_inherit` (Boolean) Use inheritence mechanism
- `bundle_coexistence_bob_blb_logical` (Boolean) Use BFD logical Bundle natively
- `dampening_bundle_member_initial_wait` (Number) Initial delay before bringing up session
  - Range: `1`-`518400000`
- `dampening_bundle_member_l3_only_mode` (Boolean) Apply immediate dampening and only when failure is L3 specific
- `dampening_bundle_member_maximum_wait` (Number) Maximum delay before bringing up session
  - Range: `1`-`518400000`
- `dampening_bundle_member_secondary_wait` (Number) Secondary delay before bringing up session
  - Range: `1`-`518400000`
- `dampening_disable` (Boolean) Disable BFD dampening
- `dampening_extensions_down_monitoring` (Boolean) Enable DOWN state session monitoring extensions
- `dampening_initial_wait` (Number) Initial delay before bringing up session
  - Range: `1`-`3600000`
- `dampening_maximum_wait` (Number) Maximum delay before bringing up session
  - Range: `1`-`3600000`
- `dampening_secondary_wait` (Number) Secondary delay before bringing up session
  - Range: `1`-`3600000`
- `dampening_threshold` (Number) Stability threshold to enable dampening
  - Range: `60000`-`3600000`
- `delete_mode` (String) Configure behavior when deleting/destroying the resource. Either delete the entire object (YANG container) being managed, or only delete the individual resource attributes configured explicitly and leave everything else as-is. Default value is `all`.
  - Choices: `all`, `attributes`
- `device` (String) A device name from the provider configuration.
- `echo_disable` (Boolean) Disable BFD echo mode
- `echo_ipv4_bundle_per_member_preferred_minimum_interval` (Number) The preferred minimum interval (in ms) for the BFD session
  - Range: `15`-`2000`
- `echo_ipv4_source` (String) BFD echo source IP address
- `echo_latency_detect_count` (Number) Count of consecutive bad latency packets to take session down
  - Range: `1`-`10`
- `echo_latency_detect_percentage` (Number) Percentage of detection time to consider as bad latency
  - Range: `100`-`250`
- `echo_startup_validate_force` (Boolean) Ignore remote 'Required Min Echo RX Interval' setting
- `interfaces` (Attributes List) Configure BFD on an interface (see [below for nested schema](#nestedatt--interfaces))
- `ipv6_checksum_disable` (Boolean) Disable BFD checksum
- `multihop_ttl_drop_threshold` (Number) TTL Drop Threshold
  - Range: `0`-`254`
- `multipath_locations` (Attributes List) Specify a location (see [below for nested schema](#nestedatt--multipath_locations))
- `trap_singlehop_pre_mapped` (Boolean) Configure BFD trap pre-mapped

### Read-Only

- `id` (String) The path of the object.

<a id="nestedatt--interfaces"></a>
### Nested Schema for `interfaces`

Required:

- `interface_name` (String) Configure BFD on an interface

Optional:

- `disable` (Boolean) Disable BFD for this interface
- `echo_disable` (Boolean) Disable BFD echo mode for this interface
- `echo_ipv4_source` (String) BFD echo source IP address
- `ipv6_checksum_disable` (Boolean) Disable BFD ipv6 checksum mode for this interface
- `local_address` (String) Local address to be used by BFD for this interface
- `multiplier` (Number) BFD multiplier for this interface
  - Range: `2`-`50`
- `rx_interval` (Number) BFD RX Interval for this interface in microseconds
  - Range: `3000`-`30000000`
- `tx_interval` (Number) BFD TX Interval for this interface in microseconds
  - Range: `3000`-`30000000`


<a id="nestedatt--multipath_locations"></a>
### Nested Schema for `multipath_locations`

Required:

- `location_name` (String) Specify a location

## Import

Import is supported using the following syntax:

```shell
terraform import iosxr_bfd.example ""
```