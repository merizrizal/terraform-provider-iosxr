// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIosxrNTP(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "ipv4_precedence", "network"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "ipv6_dscp", "af11"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "access_group_ipv6_peer", "peer1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "access_group_ipv6_query_only", "query1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "access_group_ipv6_serve", "serve1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "access_group_ipv6_serve_only", "serve-only123"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "access_group_ipv4_peer", "peer1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "access_group_ipv4_query_only", "query1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "access_group_ipv4_serve", "serve1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "access_group_ipv4_serve_only", "serve-only123"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "access_group_vrfs.0.vrf_name", "ntp_vrf"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "access_group_vrfs.0.ipv6_peer", "peer1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "access_group_vrfs.0.ipv6_query_only", "query1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "access_group_vrfs.0.ipv6_serve", "serve1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "access_group_vrfs.0.ipv6_serve_only", "serve-only123"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "access_group_vrfs.0.ipv4_peer", "peer1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "access_group_vrfs.0.ipv4_query_only", "query1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "access_group_vrfs.0.ipv4_serve", "serve1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "access_group_vrfs.0.ipv4_serve_only", "serve-only123"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "authenticate", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "authentication_keys.0.key_number", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "authentication_keys.0.md5_encrypted", "1212000E43"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "broadcastdelay", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "max_associations", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "trusted_keys.0.key_number", "8"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "update_calendar", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "log_internal_sync", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "source_interface_name", "BVI1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "source_vrfs.0.vrf_name", "source_vrf"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "source_vrfs.0.interface_name", "BVI1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "passive", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "cmac_authentication_keys.0.key_number", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "cmac_authentication_keys.0.cmac_encrypted", "135445415F59527D737D78626771475240"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "hmac_sha1_authentication_keys.0.key_number", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "hmac_sha1_authentication_keys.0.hmac_sha1_encrypted", "101F5B4A5142445C545D7A7A767B676074"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "hmac_sha2_authentication_keys.0.key_number", "4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "hmac_sha2_authentication_keys.0.hmac_sha2_encrypted", "091D1C5A4D5041455355547B79777C6663"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "ipv4_peers_servers.0.address", "1.2.3.4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "ipv4_peers_servers.0.type", "server"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "ipv4_peers_servers.0.version", "2"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "ipv4_peers_servers.0.key", "1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "ipv4_peers_servers.0.minpoll", "4"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "ipv4_peers_servers.0.maxpoll", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "ipv4_peers_servers.0.prefer", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "ipv4_peers_servers.0.burst", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "ipv4_peers_servers.0.iburst", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_ntp.test", "ipv4_peers_servers.0.source", "GigabitEthernet0/0/0/1"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrNTPConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrNTPConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_ntp.test",
		ImportState:   true,
		ImportStateId: "Cisco-IOS-XR-um-ntp-cfg:/ntp",
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func testAccIosxrNTPConfig_minimum() string {
	config := `resource "iosxr_ntp" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrNTPConfig_all() string {
	config := `resource "iosxr_ntp" "test" {` + "\n"
	config += `	ipv4_precedence = "network"` + "\n"
	config += `	ipv6_dscp = "af11"` + "\n"
	config += `	access_group_ipv6_peer = "peer1"` + "\n"
	config += `	access_group_ipv6_query_only = "query1"` + "\n"
	config += `	access_group_ipv6_serve = "serve1"` + "\n"
	config += `	access_group_ipv6_serve_only = "serve-only123"` + "\n"
	config += `	access_group_ipv4_peer = "peer1"` + "\n"
	config += `	access_group_ipv4_query_only = "query1"` + "\n"
	config += `	access_group_ipv4_serve = "serve1"` + "\n"
	config += `	access_group_ipv4_serve_only = "serve-only123"` + "\n"
	config += `	access_group_vrfs = [{` + "\n"
	config += `		vrf_name = "ntp_vrf"` + "\n"
	config += `		ipv6_peer = "peer1"` + "\n"
	config += `		ipv6_query_only = "query1"` + "\n"
	config += `		ipv6_serve = "serve1"` + "\n"
	config += `		ipv6_serve_only = "serve-only123"` + "\n"
	config += `		ipv4_peer = "peer1"` + "\n"
	config += `		ipv4_query_only = "query1"` + "\n"
	config += `		ipv4_serve = "serve1"` + "\n"
	config += `		ipv4_serve_only = "serve-only123"` + "\n"
	config += `		}]` + "\n"
	config += `	authenticate = true` + "\n"
	config += `	authentication_keys = [{` + "\n"
	config += `		key_number = 10` + "\n"
	config += `		md5_encrypted = "1212000E43"` + "\n"
	config += `		}]` + "\n"
	config += `	broadcastdelay = 10` + "\n"
	config += `	max_associations = 1` + "\n"
	config += `	trusted_keys = [{` + "\n"
	config += `		key_number = 8` + "\n"
	config += `		}]` + "\n"
	config += `	update_calendar = true` + "\n"
	config += `	log_internal_sync = true` + "\n"
	config += `	source_interface_name = "BVI1"` + "\n"
	config += `	source_vrfs = [{` + "\n"
	config += `		vrf_name = "source_vrf"` + "\n"
	config += `		interface_name = "BVI1"` + "\n"
	config += `		}]` + "\n"
	config += `	passive = true` + "\n"
	config += `	cmac_authentication_keys = [{` + "\n"
	config += `		key_number = 2` + "\n"
	config += `		cmac_encrypted = "135445415F59527D737D78626771475240"` + "\n"
	config += `		}]` + "\n"
	config += `	hmac_sha1_authentication_keys = [{` + "\n"
	config += `		key_number = 3` + "\n"
	config += `		hmac_sha1_encrypted = "101F5B4A5142445C545D7A7A767B676074"` + "\n"
	config += `		}]` + "\n"
	config += `	hmac_sha2_authentication_keys = [{` + "\n"
	config += `		key_number = 4` + "\n"
	config += `		hmac_sha2_encrypted = "091D1C5A4D5041455355547B79777C6663"` + "\n"
	config += `		}]` + "\n"
	config += `	ipv4_peers_servers = [{` + "\n"
	config += `		address = "1.2.3.4"` + "\n"
	config += `		type = "server"` + "\n"
	config += `		version = 2` + "\n"
	config += `		key = 1` + "\n"
	config += `		minpoll = 4` + "\n"
	config += `		maxpoll = 5` + "\n"
	config += `		prefer = true` + "\n"
	config += `		burst = true` + "\n"
	config += `		iburst = true` + "\n"
	config += `		source = "GigabitEthernet0/0/0/1"` + "\n"
	config += `		}]` + "\n"
	config += `}` + "\n"
	return config
}
