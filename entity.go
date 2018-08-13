package main

import "github.com/boltdb/bolt"

// Entity defines a base set of crud functions.
type Entity interface {
	Create(db *bolt.DB) error
	Update(db *bolt.DB) error
	Delete(db *bolt.DB, state bool) error
}
