package dnstools

import (
	"errors"
	"net"
	"strings"
)


type Record struct {
	Domain     string
	Host       string
	RecordType string
}

func New(domain, RecordType string) *Record {
	return &Record{
		Domain:     domain,
		RecordType: RecordType,
	}
}

func (re *Record) getCNAME() error {
	result, err := net.LookupCNAME(re.Domain)
	if err != nil {
		return  err
	}
	if result != re.Domain + "." {
		re.Host = result
	}
	return nil
}

func (re *Record) getA() error {
	result, err := net.LookupHost(re.Domain)
	if err != nil {
		return err
	}
	re.Host = strings.Join(result, " ")
	return nil
}

func (re *Record) GetRecordInfo() error {
	switch {
	case re.RecordType == "A":
		err := re.getA()
		if err != nil {
			return err
		}
	case re.RecordType == "CNAME":
		err := re.getCNAME()
		if err != nil {
			return err
		}
	default:
		return errors.New("type error")
	}

	return nil
}


