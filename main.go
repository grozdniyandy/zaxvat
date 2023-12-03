package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net"
    "net/http"
    "os"
    "strings"
)

type Fingerprint struct {
    CICDPass       bool     `json:"cicd_pass"`
    CNAME          []string `json:"cname"`
    Discussion     string   `json:"discussion"`
    Documentation string   `json:"documentation"`
    Fingerprint    string   `json:"fingerprint"`
    HTTPStatus     int      `json:"http_status"`
    NXDomain       bool     `json:"nxdomain"`
    Service        string   `json:"service"`
    Status         string   `json:"status"`
    Vulnerable     bool     `json:"vulnerable"`
}

func main() {
    fmt.Println("Made with lyubofff by GrozdniyAndy of XSS.is")
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run main.go <domain>")
        os.Exit(1)
    }

    domain := os.Args[1]

    //fmt.Println("Fetching fingerprints...")
    fingerprints, err := getFingerprints()
    if err != nil {
        fmt.Println("Error fetching fingerprints:", err)
        os.Exit(1)
    }

    //fmt.Println("Fingerprints fetched successfully.")

    // Print the list of fetched fingerprints
    //fmt.Println("List of fetched fingerprints:")
    //for i, fp := range fingerprints {
        //fmt.Printf("%d. Service: %s, CNAME: %v\n", i+1, fp.Service, fp.CNAME)
    //}

    fmt.Printf("Checking vulnerability for domain: %s\n", domain)
    checkVulnerability(domain, fingerprints)
}

func getFingerprints() ([]Fingerprint, error) {
    url := "https://raw.githubusercontent.com/EdOverflow/can-i-take-over-xyz/master/fingerprints.json"

    // log.Printf("Fetching fingerprints from %s\n", url)
    response, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }

    var fingerprints []Fingerprint
    err = json.Unmarshal(body, &fingerprints)
    if err != nil {
        return nil, err
    }

    // log.Printf("Fingerprints fetched successfully\n")
    return fingerprints, nil
}

func checkVulnerability(domain string, fingerprints []Fingerprint) {
    cname, err := getCNAME(domain)
    if err != nil {
        log.Fatalf("Error fetching CNAME for %s: %v\n", domain, err)
    }

    fmt.Printf("CNAME for %s: %s\n", domain, cname)

    for _, fp := range fingerprints {
        for _, fpCname := range fp.CNAME {
            if strings.Contains(cname, fpCname) {
                fmt.Printf("CNAME for %s matches with a service in the list:\n", domain)
                fmt.Printf("Service: %s\n", fp.Service)
                fmt.Printf("Status: %s\n", fp.Status)
                fmt.Printf("Discussion: %s\n", fp.Discussion)
                fmt.Printf("Documentation: %s\n", fp.Documentation)
                return
            }
        }
    }

    fmt.Printf("CNAME for %s is not found in the vulnerable fingerprints.\n", domain)
}

func getCNAME(domain string) (string, error) {
    cname, err := net.LookupCNAME(domain)
    if err != nil {
        return "", err
    }
    return cname, nil
}
