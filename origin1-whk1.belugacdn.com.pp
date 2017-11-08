node 'origin-whk1.belugacdn.com' {
  class{'beluga-network':
    loopbacks  => [],
    interfaces => [
      {
        interface    => 'em1',
        gateway      => '66.230.130.93',
        address      => '66.230.130.94',
        netmask      => '255.255.255.252',
        ipv6_gateway => 'fe80::92b1:1cff:fe12:cd4',
        ipv6_address => 'fe80::92b1:1cff:fe12:cd5',
        ipv6_netmask => '126',
        default      => true
      }
    ]
  }
  include beluga_nginx
}
