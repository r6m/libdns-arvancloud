package arvancloud

import (
	"encoding/json"

	"github.com/libdns/libdns"
)

type DnsRecordList struct {
	Data []DnsRecord
}

type DnsRecord struct {
	ID        string             `json:"id,omitempty"`
	Type      string             `json:"type,omitempty"`
	Name      string             `json:"name,omitempty"`
	ValueData DnsRecordValueData `json:"value,omitempty"`
	TTL       int                `json:"ttl,omitempty"`

	Value string
}

type DnsRecordValueData struct {
	IP       string `json:"string,omitempty"`
	Host     string `json:"host,omitempty"`
	Target   string `json:"target,omitempty"`
	Text     string `json:"text,omitempty"`
	Location string `json:"location,omitempty"`
	Domain   string `json:"domain,omitempty"`
}

func (v DnsRecord) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	switch v.Type {
	case "a":
		v.Value = v.ValueData.IP
	case "aaaa":
		v.Value = v.ValueData.IP
	case "mx":
		v.Value = v.ValueData.Host
	case "ns":
		v.Value = v.ValueData.Host
	case "srv":
		v.Value = v.ValueData.Target
	case "txt":
		v.Value = v.ValueData.Text
	case "spf":
		v.Value = v.ValueData.Text
	case "dkim":
		v.Value = v.ValueData.Text
	case "aname":
		v.Value = v.ValueData.Location
	case "cname":
		v.Value = v.ValueData.Host
	case "ptr":
		v.Value = v.ValueData.Domain
	}

	return nil
}

func (v DnsRecord) MarshalJSON() ([]byte, error) {
	switch v.Type {
	case "a":
		v.ValueData.IP = v.Value
	case "aaaa":
		v.ValueData.IP = v.Value
	case "mx":
		v.ValueData.Host = v.Value
	case "ns":
		v.ValueData.Host = v.Value
	case "srv":
		v.ValueData.Target = v.Value
	case "txt":
		v.ValueData.Text = v.Value
	case "spf":
		v.ValueData.Text = v.Value
	case "dkim":
		v.ValueData.Text = v.Value
	case "aname":
		v.ValueData.Location = v.Value
	case "cname":
		v.ValueData.Host = v.Value
	case "ptr":
		v.ValueData.Domain = v.Value
	}

	return json.Marshal(v)
}

func fromLibdns(record libdns.Record) DnsRecord {
	ttl := int(record.TTL.Seconds())
	if ttl == 0 {
		ttl = 120
	}

	return DnsRecord{
		ID:    record.ID,
		Type:  record.Type,
		Name:  record.Name,
		TTL:   ttl,
		Value: record.Value,
	}
}
