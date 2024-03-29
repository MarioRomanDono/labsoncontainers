// SPDX-FileCopyrightText: 2022 Mario Román Dono <mario.romandono@gmail.com>
//
// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"flag"
	"fmt"
	"os"
)

func Usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Uso: %s [opción] [nombreLab | nombreLab.yaml]\n", os.Args[0])
	fmt.Fprintln(flag.CommandLine.Output(), "")
	fmt.Fprintln(flag.CommandLine.Output(), "Opciones:")
	flag.PrintDefaults()
}

func exitOnError(err string) {
	fmt.Fprintf(flag.CommandLine.Output(), "%s\n", err)
	flag.CommandLine.Usage()
	os.Exit(1)
}

func main() {
	flag.CommandLine.Usage = Usage

	create := flag.Bool("c", false, "Crea el entorno de contenedores a partir del fichero YAML proporcionado (se destruyen los contenedores asociados al entorno del fichero)")
	start := flag.Bool("l", false, "Ejecuta todos los contenedores asociados al entorno proporcionado")
	inspect := flag.Bool("i", false, "Muestra la información de todos los contenedores asociados al entorno proporcionado")
	stop := flag.Bool("p", false, "Detiene todos los contenedores asociados al entorno proporcionado")
	destroy := flag.Bool("r", false, "Destruye todos los contenedores asociados al entorno proporcionado")

	flag.Parse()

	if flag.NFlag() < 1 {
		exitOnError("Debe proporcionar una opción válida")
	} else if flag.NFlag() > 1 {
		exitOnError("Debe proporcionar una sola opción")
	} else if flag.NArg() < 1 {
		exitOnError("Debe proporcionar un nombre de archivo o un nombre de entorno")
	} else if flag.NArg() > 1 {
		exitOnError("Debe proporcionar un solo nombre de archivo o nombre de entorno")
	}

	if *create {
		createLabEnvironment(flag.Arg(0))
	} else if *start {
		startLabEnvironment(flag.Arg(0))
	} else if *inspect {
		inspectLabEnvironment(flag.Arg(0))
	} else if *stop {
		stopLabEnvironment(flag.Arg(0))
	} else if *destroy {
		destroyLabEnvironment(flag.Arg(0))
	} else {
		fmt.Fprintln(flag.CommandLine.Output(), "Opción no implementada")
		os.Exit(1)
	}
}
