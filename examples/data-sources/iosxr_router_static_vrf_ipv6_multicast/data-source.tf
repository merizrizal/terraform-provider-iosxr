data "iosxr_router_static_vrf_ipv6_multicast" "example" {
  vrf_name       = "VRF2"
  prefix_address = "1::"
  prefix_length  = 64
}
