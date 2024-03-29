// SPDX-FileCopyrightText: 2022 Mario Román Dono <mario.romandono@gmail.com>
//
// SPDX-License-Identifier: GPL-3.0-or-later

// Package labsoncontainers allows to create and manage lab environments using Docker Engine API. These environments are designed
// to practice Computer Science exercises: monitor network traffic, create VPNs, perform mock cyberattacks...
//
// The created containers have the following features:
//
// • Allow interaction with the container (-it)
//
// • Management of network settings (--cap-add=NET_ADMIN)
//
// • Running init system (--init)
//
// • Execution of GUI apps is possible using X11 untrusted cookies (https://github.com/mviereck/x11docker/wiki/X-authentication-with-cookies-and-xhost-("No-protocol-specified"-error)#untrusted-cookie-for-container-applications)
//
// • Home directory is bind mounted on all the lab containers
package labsoncontainers

// LabEnvironment represents the structure of a lab environment. It is composed by the lab
// name and a list of LabContainer.
type LabEnvironment struct {
	LabName    string         `yaml:"nombre_practica"`
	Containers []LabContainer `yaml:"contenedores"`
}

// LabContainer represents the structure of a lab container. It is composed by its name, the image name
// the container will use, a list of LabNetwork and the background field: if it is set to true, a terminal
// window will not be created for interacting with the container.
type LabContainer struct {
	Name       string       `yaml:"nombre"`
	Image      string       `yaml:"imagen"`
	Networks   []LabNetwork `yaml:"redes"`
	Background bool         `yaml:"background,omitempty"`
	ID         string
}

// LabNetwork represents the structure of a lab network. It is composed by its name and, optionally,
// the IP that will be used by the container in the network.
type LabNetwork struct {
	Name string `yaml:"nombre"`
	IP   string `yaml:"ip"`
}

// cookieDir sets where X11 cookies will be stored
const cookieDir string = "/var/tmp/X11Cookies/"
