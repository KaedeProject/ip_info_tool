package main

import (
	"bufio"
	"encoding/json"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/KaedeProject/ip_info_tool/logger"
)

type IPInfo struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Org      string `json:"org"`
	Loc      string `json:"loc"`
	Timezone string `json:"timezone"`
}

func main() {
	logger.Log("Please enter an IP address or domain: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	logger.Log("\n\nThe tool has started...\n")

	logger.Log("Getting IP...")
	ipaddr, err := net.LookupIP(input)
	if err != nil {
		logger.Log("Error getting IP: " + err.Error())
		return
	}
	ip := ipaddr[0].String()
	logger.LogVIp(ip)

	logger.Log("Getting location info...")
	req, err := http.Get("https://ipinfo.io/" + ip + "/json")
	if err != nil {
		logger.Log("Error getting location info: " + err.Error())
		return
	}
	defer req.Body.Close()

	var ipinfo IPInfo
	if err := json.NewDecoder(req.Body).Decode(&ipinfo); err != nil {
		logger.Log("Error decoding location info: " + err.Error())
		return
	}

	logger.Log("\n\nLocation info:")
	logger.LogVIp(ipinfo.IP)
	logger.LogVCity(ipinfo.City)
	logger.LogVnRegion(ipinfo.Region)
	logger.LogVCountry(ipinfo.Country)
	logger.LogVOrg(ipinfo.Org)

	logger.Log("Press Enter to exit...")
	reader.ReadBytes('\n')

}
