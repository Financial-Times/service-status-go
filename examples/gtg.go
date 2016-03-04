package main

import (
	"errors"
	"github.com/Financial-Times/service-status-go/gtg"
)

func main() {
        statusChecker = myServiceIsBroken()
        status = gtg.Status{}
}

func myServiceIsBroken() error {
	return errors.New("Some error occured")


        v1a.Handler("PublicBrandsRead Healthchecks", "Checks for accessing neo4j", brands.HealthCheck()
}
