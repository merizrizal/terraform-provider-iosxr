---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_flow_exporter_map Data Source - terraform-provider-iosxr"
subcategory: "Flow"
description: |-
  This data source can read the Flow Exporter Map configuration.
---

# iosxr_flow_exporter_map (Data Source)

This data source can read the Flow Exporter Map configuration.

## Example Usage

```terraform
data "iosxr_flow_exporter_map" "example" {
  name = "TEST"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Exporter map name - maximum 32 characters

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `destination_ipv4_address` (String) Destination IPv4 address
- `destination_ipv6_address` (String) Destination IPv6 address
- `destination_vrf` (String) Configure VRF to be used for reaching export destination
- `dfbit_set` (Boolean) Set Export Packet Do Not Fragment Flag
- `dscp` (Number) Specify DSCP value for ipv4 export packets or traffic-class for ipv6 export packets
- `id` (String) The path of the retrieved object.
- `packet_length` (Number) Export Packet maximum L3 length, should conform to outgoing interface mtu
- `source` (String) Source interface
- `transport_udp` (Number) Use UDP as transport protocol
- `version_export_format` (String) Specify export format
- `version_options_class_table_timeout` (Number) Specify export timeout
- `version_options_interface_table_timeout` (Number) Specify export timeout
- `version_options_sampler_table_timeout` (Number) Specify export timeout
- `version_options_vrf_table_timeout` (Number) Specify export timeout
- `version_template_data_timeout` (Number) Specify custom timeout for the template
- `version_template_options_timeout` (Number) Specify custom timeout for the template
- `version_template_timeout` (Number) Specify custom timeout for the template