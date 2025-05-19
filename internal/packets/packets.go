package packets

import (
    "github.com/google/gopacket/pcap"
)

func GetAllDevices() (*[]pcap.Interface, error) {
    devices, err := pcap.FindAllDevs()
    if err != nil {
        return nil, err 
    }

    return &devices, nil
}
