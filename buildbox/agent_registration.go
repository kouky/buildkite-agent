package buildbox

import (
	"fmt"
	"github.com/Sirupsen/logrus"
)

type AgentRegistration struct {
	// The access token for the agent
	AccessToken string `json:"access_token"`

	// Hostname of the machine
	Hostname string `json:"hostname"`

	// The name of the new agent
	Name string `json:"name"`

	// Meta data for the agent
	MetaData []string `json:"meta_data"`
}

func (c *Client) AgentRegister(name string, metaData []string) (string, error) {
	// Create the agent registration
	var registration AgentRegistration
	registration.Name = name
	registration.Hostname = MachineHostname()
	registration.MetaData = metaData

	Logger.WithFields(logrus.Fields{
		"name":      registration.Name,
		"hostname":  registration.Hostname,
		"meta-data": registration.MetaData,
	}).Info("Registering agent")

	// Register and return the agent
	err := c.Post(&registration, "/register", registration)
	if err != nil {
		return "", err
	}

	return registration.AccessToken, nil
}

func (a *AgentRegistration) String() string {
	return fmt.Sprintf("AgentRegistration{AccessToken: %s}", a.AccessToken)
}
