package main

import "time"

type RoutingTable struct {
	table   map[string]string
	log     map[time.Time]string
	members map[string]bool
}

type RoutingInterface interface {
	addEntry(source string, file string)
	removeEntry(file string)
	getLocation(file string) string
	addMember(m string)
	removeMember(m string)
}

func (r RoutingTable) addEntry(source string, file string) {
	r.table[file] = source
}

func (r RoutingTable) removeEntry(file string) {
	delete(r.table, file)
}

func (r RoutingTable) getLocation(file string) string {
	return r.table[file]
}

func (r RoutingTable) addMember(m string) {
	r.members[m] = true
}

func (r RoutingTable) removeMember(m string) {
	delete(r.members, m)
}
