//go:build !no_cassandra
// +build !no_cassandra

package main

import (
	_ "github.com/gocql/gocql"
)
