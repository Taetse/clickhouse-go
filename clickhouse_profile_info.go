package clickhouse

type ProfileInfo struct {
	Rows                      uint64
	Bytes                     uint64
	Blocks                    uint64
	AppliedLimit              bool
	RowsBeforeLimit           uint64
	CalculatedRowsBeforeLimit bool
}

func (ch *clickhouse) profileInfo() (*ProfileInfo, error) {
	var (
		p   ProfileInfo
		err error
	)
	if p.Rows, err = ch.decoder.Uvarint(); err != nil {
		return nil, err
	}
	if p.Blocks, err = ch.decoder.Uvarint(); err != nil {
		return nil, err
	}
	if p.Bytes, err = ch.decoder.Uvarint(); err != nil {
		return nil, err
	}

	if p.AppliedLimit, err = ch.decoder.Bool(); err != nil {
		return nil, err
	}
	if p.RowsBeforeLimit, err = ch.decoder.Uvarint(); err != nil {
		return nil, err
	}
	if p.CalculatedRowsBeforeLimit, err = ch.decoder.Bool(); err != nil {
		return nil, err
	}
	return &p, nil
}
