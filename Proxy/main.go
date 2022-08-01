package main

import (
	"fmt"
	"time"
)

type server interface {
	handleRequest(string) (int, string)
}

type Licence struct {
}

/*if licence is available at the moment, return success*/
func (a *Licence) handleRequest(request string) (int, string) {

	/*Licences*/
	if request == "VLSI" {
		fmt.Printf("\n request: %s\nHttpCode: %d\nBody: %s\n", request, 200, "VLSI Provided")
		return 200, "VLSI Provided"
	}

	if request == "Operating System - Linux" {
		fmt.Printf("\n request: %s\nHttpCode: %d\nBody: %s\n", request, 200, "Linux Provided")
		return 200, "Linux Provided"
	}

	if request == "Virtual CPU" {
		fmt.Printf("\n request: %s\nHttpCode: %d\nBody: %s\n", request, 200, "Virtual CPU Provided")
		return 200, "Virtual CPU Provided"
	}

	fmt.Printf("\n request: %s\nHttpCode: %d\nBody: %s\n", request, 404, "Source Cannot Found")
	return 404, "Source Cannot Found"
}

type LoadBalancer struct {
	licence           *Licence
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func newLoadBalancer() *LoadBalancer {
	return &LoadBalancer{
		licence:           &Licence{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}

}

func (n *LoadBalancer) handleRequest(request string) (int, string) {
	allowed := n.checkRateLimiting(request)
	if !allowed {
		fmt.Printf("\n request: %s\nHttpCode: %d\nBody: %s\n", request, 403, "Not Allowed")
		return 403, "Not Allowed"
	}
	return n.licence.handleRequest(request)
}

func (n *LoadBalancer) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}

func main() {

	LoadBalancer := newLoadBalancer()
	neededLicence1 := "VLSI"
	neededLicence2 := "Operating System - Linux"
	neededLicence3 := "Virtual CPU"
	neededLicence4 := "VLSI2"

	for i := 0; i < 3; i++ {

		go LoadBalancer.handleRequest(neededLicence1)
		time.Sleep(100 * time.Millisecond)

		go LoadBalancer.handleRequest(neededLicence2)
		time.Sleep(100 * time.Millisecond)

		go LoadBalancer.handleRequest(neededLicence3)
		time.Sleep(100 * time.Millisecond)

		go LoadBalancer.handleRequest(neededLicence4)
		time.Sleep(100 * time.Millisecond)
	}

}
