// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// +godefs map struct_in6_addr [16]byte /* in6_addr */

package ipv6

/*
#define __APPLE_USE_RFC_3542
#include <netinet/in.h>
#include <netinet/icmp6.h>
*/
import "C"

const (
	sysIPV6_TCLASS   = C.IPV6_TCLASS
	sysIPV6_PATHMTU  = C.IPV6_PATHMTU
	sysIPV6_PKTINFO  = C.IPV6_PKTINFO
	sysIPV6_HOPLIMIT = C.IPV6_HOPLIMIT
	sysIPV6_NEXTHOP  = C.IPV6_NEXTHOP

	sizeofSockaddrStorage = C.sizeof_struct_sockaddr_storage
	sizeofSockaddrInet6   = C.sizeof_struct_sockaddr_in6
	sizeofInet6Pktinfo    = C.sizeof_struct_in6_pktinfo
	sizeofIPv6Mtuinfo     = C.sizeof_struct_ip6_mtuinfo

	sizeofIPv6Mreq       = C.sizeof_struct_ipv6_mreq
	sizeofGroupReq       = C.sizeof_struct_group_req
	sizeofGroupSourceReq = C.sizeof_struct_group_source_req

	sizeofICMPv6Filter = C.sizeof_struct_icmp6_filter
)

type sockaddrStorage C.struct_sockaddr_storage

type sockaddrInet6 C.struct_sockaddr_in6

type inet6Pktinfo C.struct_in6_pktinfo

type ipv6Mtuinfo C.struct_ip6_mtuinfo

type ipv6Mreq C.struct_ipv6_mreq

type icmpv6Filter C.struct_icmp6_filter

type groupReq C.struct_group_req

type groupSourceReq C.struct_group_source_req
