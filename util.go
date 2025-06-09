package cidr_util

import (
	"crypto/rand"
	"net"
	"fmt"
)

type IPCidr struct {
	base net.IP
	mask net.IPMask
}

func ParseCIDR(cidr string) (*IPCidr, error) {
	ip, ip_net, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	return &IPCidr{base: ip, mask: ip_net.Mask}, nil
}

func (c *IPCidr) RandomIP() net.IP {
	var ipv int

	if c.base.To4() != nil {
		ipv = 4
	} else if c.base.To16() != nil {
		ipv = 16
	} else {
		return nil
	}

	mask_size, _ := c.mask.Size()
	bits := ipv * 8
	randbits := bits - mask_size
	randbytes := (randbits + 7) / 8

	ip := make([]byte, ipv)
	copy(ip, c.base)

	r := make([]byte, randbytes)
	if _, err := rand.Read(r); err != nil {
		return nil
	}

	for i := range randbytes {
		idx := mask_size/8 + i
		if idx >= ipv {
			break
		}
		if i == 0 && mask_size%8 != 0 {
			mask := byte(0xFF >> (mask_size % 8))
			ip[idx] = ip[idx]&^mask | (r[i] & mask)
		} else {
			ip[idx] = r[i]
		}
	}

	return net.IP(ip)
}

func (c *IPCidr) String() string {
	mask_size, _ := c.mask.Size()
	return fmt.Sprintf("%s/%d", c.base.String(), mask_size)
}
