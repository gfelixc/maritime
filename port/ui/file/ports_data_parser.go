package file

import (
	"bufio"
	"encoding/json"
	"io"
)

type PortsDataParser struct {
	source *bufio.Reader
}

func NewPortsDataParser(source io.Reader) *PortsDataParser {
	return &PortsDataParser{source: bufio.NewReader(source)}
}

func (p *PortsDataParser) NextPort() (Port, error) {
	fieldName, err := p.getFieldName()
	if err != nil {
		return Port{}, err
	}

	object, err := p.getObject()
	if err != nil {
		return Port{}, err
	}

	port := Port{}
	err = json.Unmarshal([]byte(object), &port)
	if err != nil {
		panic(err)
	}

	port.FieldName = fieldName

	return port, err
}

func (p *PortsDataParser) getFieldName() (string, error) {
	_, err := p.source.ReadString('"')
	if err != nil {
		return "", err
	}

	id, err := p.source.ReadString('"')
	if err != nil {
		return "", err
	}

	return id[:len(id)-1], nil
}

func (p *PortsDataParser) getObject() (string, error) {
	_, err := p.source.ReadString('{')
	if err != nil {
		return "", err
	}

	chunk, err := p.source.ReadString('}')
	if err != nil {
		return "", err
	}

	return "{" + chunk, nil
}

type Port struct {
	FieldName   string
	Name        string
	City        string
	Country     string
	Alias       []string
	Regions     []string
	Coordinates []float64
	Province    *string
	Timezone    *string
	Unlocs      []string
	Code        *string
}
