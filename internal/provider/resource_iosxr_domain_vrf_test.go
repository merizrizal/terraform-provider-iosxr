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

func TestAccIosxrDomainVRF(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "vrf_name", "TEST-VRF"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "domains.0.domain_name", "DOMAIN11"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "domains.0.order", "12345"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "lookup_disable", "true"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "lookup_source_interface", "Loopback2147483647"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "name", "DNAME"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "ipv4_hosts.0.host_name", "HOST-AGC"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "ipv4_hosts.0.ip_address.0", "10.0.0.0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "name_servers.0.address", "10.0.0.1"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "name_servers.0.order", "0"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "ipv6_hosts.0.host_name", "HOST-ACC"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "ipv6_hosts.0.ipv6_address.0", "10::10"))
	checks = append(checks, resource.TestCheckResourceAttr("iosxr_domain_vrf.test", "multicast", "TESTACC"))
	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccIosxrDomainVRFConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccIosxrDomainVRFConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})
	steps = append(steps, resource.TestStep{
		ResourceName:  "iosxr_domain_vrf.test",
		ImportState:   true,
		ImportStateId: "TEST-VRF",
		Check:         resource.ComposeTestCheckFunc(checks...),
	})
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

func testAccIosxrDomainVRFConfig_minimum() string {
	config := `resource "iosxr_domain_vrf" "test" {` + "\n"
	config += `	vrf_name = "TEST-VRF"` + "\n"
	config += `}` + "\n"
	return config
}

func testAccIosxrDomainVRFConfig_all() string {
	config := `resource "iosxr_domain_vrf" "test" {` + "\n"
	config += `	vrf_name = "TEST-VRF"` + "\n"
	config += `	domains = [{` + "\n"
	config += `		domain_name = "DOMAIN11"` + "\n"
	config += `		order = 12345` + "\n"
	config += `		}]` + "\n"
	config += `	lookup_disable = true` + "\n"
	config += `	lookup_source_interface = "Loopback2147483647"` + "\n"
	config += `	name = "DNAME"` + "\n"
	config += `	ipv4_hosts = [{` + "\n"
	config += `		host_name = "HOST-AGC"` + "\n"
	config += `		ip_address = ["10.0.0.0"]` + "\n"
	config += `		}]` + "\n"
	config += `	name_servers = [{` + "\n"
	config += `		address = "10.0.0.1"` + "\n"
	config += `		order = 0` + "\n"
	config += `		}]` + "\n"
	config += `	ipv6_hosts = [{` + "\n"
	config += `		host_name = "HOST-ACC"` + "\n"
	config += `		ipv6_address = ["10::10"]` + "\n"
	config += `		}]` + "\n"
	config += `	multicast = "TESTACC"` + "\n"
	config += `}` + "\n"
	return config
}