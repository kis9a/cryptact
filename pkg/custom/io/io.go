package io

import (
	"errors"
	"fmt"
	"io"
	"math/big"
	"time"

	"github.com/kis9a/cryptact/pkg/custom"
)

type CustomIO struct {
	DateFormat string
}

type CustomIoIF interface {
	Read(reader io.Reader) ([]custom.CustomData, error)
	Write(writer io.Writer, data []custom.CustomData) error
}

type CustomIoData struct {
	Timestamp string `csv:"timestamp"`
	Action    string `csv:"action"`
	Source    string `csv:"source"`
	Base      string `csv:"base"`
	Volume    string `csv:"volume"`
	Price     string `csv:"price"`
	Counter   string `csv:"counter"`
	Fee       string `csv:"fee"`
	FeeCcy    string `csv:"feeccy"`
	Comment   string `csv:"comment"`
}

func (c *CustomIO) NewIoInstance() {
	c.DateFormat = "2006/01/02 15:04:05"
}

func (c *CustomIO) ParseCustomIoDataToCustomData(record CustomIoData) (custom.CustomData, error) {
	var customData custom.CustomData
	timestamp, err := time.Parse(c.DateFormat, record.Timestamp)
	if err != nil {
		return customData, err
	}
	volume, ok := new(big.Float).SetString(record.Volume)
	if !ok {
		return customData, errors.New(fmt.Sprintf("Invalid type column `volume`: %s\n", record.Volume))
	}
	price, ok := new(big.Float).SetString(record.Price)
	if !ok {
		return customData, errors.New(fmt.Sprintf("Invalid type column `price`: %s\n", record.Price))
	}
	fee, ok := new(big.Float).SetString(record.Fee)
	if !ok {
		return customData, errors.New(fmt.Sprintf("Invalid type column `fee`: %s\n", record.Fee))
	}
	return custom.CustomData{
		Timestamp: timestamp,
		Action:    custom.Action(record.Action),
		Source:    record.Source,
		Base:      record.Base,
		Volume:    volume,
		Price:     price,
		Counter:   record.Counter,
		Fee:       fee,
		FeeCcy:    record.FeeCcy,
		Comment:   record.Comment,
	}, err
}

func (c *CustomIO) FormatCustomDataToCustomIoData(record custom.CustomData) CustomIoData {
	var customIoData CustomIoData
	customIoData = CustomIoData{
		Timestamp: record.Timestamp.Format(c.DateFormat),
		Action:    string(record.Action),
		Source:    record.Source,
		Base:      record.Base,
		Counter:   record.Counter,
		FeeCcy:    record.FeeCcy,
		Comment:   record.Comment,
	}
	if record.Volume != nil {
		customIoData.Volume = record.Volume.String()
	}
	if record.Price != nil {
		customIoData.Price = record.Price.String()
	}
	if record.Fee != nil {
		customIoData.Fee = record.Fee.String()
	}
	return customIoData
}
