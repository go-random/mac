package mac

import (
	"fmt"
	"math/rand"
)

// MAC generates a random MAC address.
func (tr *Randomizer) MAC() string {
	// Generate six random octets
	octets := generateSixRandomOctets(tr.Rand)

	// Format the MAC address with colons
	macAddress := fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", octets[0], octets[1], octets[2], octets[3], octets[4], octets[5])

	return macAddress
}

// GlobalMAC generates a random globally unique MAC address.
func (tr *Randomizer) GlobalMAC() string {
	// Generate six random octets
	octets := generateSixRandomOctets(tr.Rand)

	// Set the second least significant bit of the first octet to 0 to indicate it's a globally unique MAC address
	octets[0] &^= 0x02

	// Format the MAC address with colons
	macAddress := fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", octets[0], octets[1], octets[2], octets[3], octets[4], octets[5])

	return macAddress
}

// LocalMAC generates a random locally administered MAC address.
func (tr *Randomizer) LocalMAC() string {
	// Generate six random octets
	octets := generateSixRandomOctets(tr.Rand)

	// Set the second least significant bit of the first octet to 1 to indicate it's a locally administered MAC address
	octets[0] |= 0x02

	// Format the MAC address with colons
	macAddress := fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", octets[0], octets[1], octets[2], octets[3], octets[4], octets[5])

	return macAddress
}

func generateSixRandomOctets(randomizer *rand.Rand) []byte {
	octets := make([]byte, 6)
	for i := range octets {
		octets[i] = byte(randomizer.Intn(256))
	}
	return octets
}
