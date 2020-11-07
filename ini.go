package main

import (
	"gopkg.in/ini.v1"
)

type KVCPair struct {
	Key     string
	Value   string
	Comment string
}

type Section struct {
	Name    string
	Comment string
	Entries []KVCPair
}

type IniDocument struct {
	Sections []Section
	Filename string
}

func readIniFile(path string) (*IniDocument, error) {
	document := new(IniDocument)
	document.Filename = path

	cfg, err := ini.Load(path)
	if err != nil {
		return nil, err
	}
	for _, section := range cfg.Sections() {
		pairs := createKVPs(section)
		document.Sections = append(document.Sections,
			Section{
				Name:    section.Name(),
				Comment: section.Comment,
				Entries: pairs,
			},
		)
	}

	return document, nil
}

func createKVPs(section *ini.Section) []KVCPair {
	var pairs []KVCPair
	for _, key := range section.Keys() {
		pair := KVCPair{
			Key:     key.Name(),
			Value:   key.Value(),
			Comment: key.Comment,
		}
		pairs = append(pairs, pair)
	}
	return pairs
}
