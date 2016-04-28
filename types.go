package vsolver

import "fmt"

type ProjectIdentifier struct {
	LocalName   ProjectName
	NetworkName string
}

func (i ProjectIdentifier) less(j ProjectIdentifier) bool {
	if i.LocalName < j.LocalName {
		return true
	}
	if j.LocalName < i.LocalName {
		return false
	}

	return i.NetworkName < j.NetworkName
}

func (i ProjectIdentifier) eq(j ProjectIdentifier) bool {
	if i.LocalName != j.LocalName {
		return false
	}
	if i.NetworkName == j.NetworkName {
		return true
	}

	if (i.NetworkName == "" && j.NetworkName == string(j.LocalName)) ||
		(j.NetworkName == "" && i.NetworkName == string(i.LocalName)) {
		return true
	}

	return false
}

func (i ProjectIdentifier) netName() string {
	if i.NetworkName == "" {
		return string(i.LocalName)
	}
	return i.NetworkName
}

func (i ProjectIdentifier) errString() string {
	if i.NetworkName == "" || i.NetworkName == string(i.LocalName) {
		return string(i.LocalName)
	}
	return fmt.Sprintf("%s (from %s)", i.LocalName, i.NetworkName)
}

type ProjectName string

type ProjectAtom struct {
	Name    ProjectName // TODO to ProjectIdentifier
	Version Version
}

var emptyProjectAtom ProjectAtom

type ProjectDep struct {
	Ident      ProjectIdentifier
	Constraint Constraint
}

type Dependency struct {
	Depender ProjectAtom
	Dep      ProjectDep
}

// ProjectInfo holds the spec and lock information for a given ProjectAtom
type ProjectInfo struct {
	pa ProjectAtom
	Manifest
	Lock
}
