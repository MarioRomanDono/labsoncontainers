// SPDX-FileCopyrightText: 2022 Mario Román Dono <mario.romandono@gmail.com>
//
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"fmt"
	"os"

	"github.com/marioromandono/labsoncontainers"
)

// inspectLabEnvironment prints to the standard output the result of running inspect 
// on the lab environment containers, using LabsOnContainers API.
func inspectLabEnvironment(labName string) {
	fmt.Printf("Información de los contenedores de %v:\n", labName)
	containersIds, err := labsoncontainers.GetEnvironmentContainers(labName)
	if err != nil {
		fmt.Println(err)
        os.Exit(1)
	}

	if len(containersIds) > 0 {
		inspectMap, err := labsoncontainers.InspectEnvironment(labName)
   		if err != nil {
			fmt.Println(err)
			os.Exit(1)
   		}
		
		for container, inspect := range inspectMap {
			fmt.Println()
			fmt.Printf("%v:\n", container)
			fmt.Println(string(inspect))
		}			
	} else {
		fmt.Println("No existen contenedores asociados a", labName)
	}
}