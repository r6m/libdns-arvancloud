package arvancloud

import (
	"context"
	"fmt"

	"github.com/libdns/libdns"
)

type Provider struct {
	APIToken string `json:"api_token,omitempty"`
}

func (p *Provider) GetRecords(ctx context.Context, zone string) ([]libdns.Record, error) {
	return nil, fmt.Errorf("TODO: not implemented")
}

func (p *Provider) AppendRecords(ctx context.Context, zone string, records []libdns.Record) ([]libdns.Record, error) {
	return nil, fmt.Errorf("TODO: not implemented")
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
