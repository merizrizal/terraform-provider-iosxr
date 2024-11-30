---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "iosxr_pce Resource - terraform-provider-iosxr"
subcategory: "System"
description: |-
  This resource can manage the PCE configuration.
---

# iosxr_pce (Resource)

This resource can manage the PCE configuration.

## Example Usage

```terraform
resource "iosxr_pce" "example" {
  address_ipv4 = "77.77.77.1"
  state_sync_ipv4s = [
    {
      address = "100.100.100.11"
    }
  ]
  peer_filter_ipv4_access_list = "Accesslist1"
  api_authentication_digest    = true
  api_sibling_ipv4             = "100.100.100.2"
  api_users = [
    {
      user_name          = "rest-user"
      password_encrypted = "00141215174C04140B"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `address_ipv4` (String) IPv4 address
- `address_ipv6` (String) IPv6 address
- `api_authentication_digest` (Boolean) Use HTTP Digest authentication (MD5)
- `api_sibling_ipv4` (String) IPv4 address of the PCE sibling
- `api_users` (Attributes List) Northbound API username (see [below for nested schema](#nestedatt--api_users))
- `delete_mode` (String) Configure behavior when deleting/destroying the resource. Either delete the entire object (YANG container) being managed, or only delete the individual resource attributes configured explicitly and leave everything else as-is. Default value is `all`.
  - Choices: `all`, `attributes`
- `device` (String) A device name from the provider configuration.
- `peer_filter_ipv4_access_list` (String) Access-list for IPv4 peer filtering
- `state_sync_ipv4s` (Attributes List) IPv4 address (see [below for nested schema](#nestedatt--state_sync_ipv4s))

### Read-Only

- `id` (String) The path of the object.

<a id="nestedatt--api_users"></a>
### Nested Schema for `api_users`

Required:

- `user_name` (String) Northbound API username

Optional:

- `password_encrypted` (String) Specify unencrypted password


<a id="nestedatt--state_sync_ipv4s"></a>
### Nested Schema for `state_sync_ipv4s`

Required:

- `address` (String) IPv4 address

## Import

Import is supported using the following syntax:

```shell
terraform import iosxr_pce.example ""
```