package arvancloud

import (
	"context"
	"fmt"
	"sync"

	"github.com/libdns/libdns"
)

type Provider struct {
	mutex    sync.Mutex
	APIToken string `json:"api_token,omitempty"`
}

func (p *Provider) GetRecords(ctx context.Context, zone string) ([]libdns.Record, error) {
	records, err := p.getRecords(ctx, zone)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (p *Provider) AppendRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {
	var appended []libdns.Record

	for _, record := range records {
		err := p.appendRecord(ctx, zone, record)
		if err != nil {
			return appended, err
		}
	}

	return appended, nil
}

func (p *Provider) SetRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {
	return nil, fmt.Errorf("TODO: not implemented")
}

func (p *Provider) DeleteRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {
	return nil, fmt.Errorf("TODO: not implemented")
}

var (
	_ libdns.RecordGetter   = (*Provider)(nil)
	_ libdns.RecordAppender = (*Provider)(nil)
	_ libdns.RecordSetter   = (*Provider)(nil)
	_ libdns.RecordDeleter  = (*Provider)(nil)
)
