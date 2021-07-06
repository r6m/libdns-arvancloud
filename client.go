package arvancloud

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/libdns/libdns"
)

var (
	baseURL = "https://napi.arvancloud.com/cdn/4.0"
)

func (p *Provider) getRecords(ctx context.Context, zone string) ([]libdns.Record, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	url := fmt.Sprintf("%s/domains/%s/dns-records", baseURL, zone)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", p.APIToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("failed with status code %d", resp.StatusCode)
	}

	var recordList DnsRecordList
	err = json.NewDecoder(resp.Body).Decode(&recordList)
	if err != nil {
		return nil, err
	}

	var records []libdns.Record

	for _, item := range recordList.Data {
		records = append(records, libdns.Record{
			ID:    item.ID,
			Type:  item.Type,
			Name:  item.Name,
			Value: item.Value,
			TTL:   time.Duration(item.TTL) * time.Second,
		})
	}

	return records, nil
}

func (p *Provider) appendRecord(ctx context.Context, zone string, record libdns.Record) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	url := fmt.Sprintf("%s/domains/%s/dns-records", baseURL, zone)

	dnsRecord := fromLibdns(record)

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(dnsRecord)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, buf)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", p.APIToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("failed with status code %d", resp.StatusCode)
	}

	return nil
}

func (p *Provider) setRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {

	return []libdns.Record{}, nil
}

func (p *Provider) deleteRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {

	return []libdns.Record{}, nil
}
