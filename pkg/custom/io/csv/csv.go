package csv

import (
	"github.com/gocarina/gocsv"
	"github.com/kis9a/cryptact/pkg/custom"
	customIo "github.com/kis9a/cryptact/pkg/custom/io"
	"io"
	"strings"
)

type CsvIO struct {
	customIo.CustomIO
}

func NewCsvIO() *CsvIO {
	c := new(CsvIO)
	c.NewIoInstance()
	return c
}

func init() {
	gocsv.SetHeaderNormalizer(strings.ToLower)
}

func (c *CsvIO) Read(reader io.Reader) ([]custom.CustomData, error) {
	var customIoData []customIo.CustomIoData
	var customData []custom.CustomData
	csvReader := gocsv.DefaultCSVReader(reader)
	if err := gocsv.UnmarshalCSV(csvReader, &customIoData); err != nil {
		return customData, err
	}
	for _, record := range customIoData {
		customDatarecord, err := c.ParseCustomIoDataToCustomData(record)
		if err != nil {
			return customData, err
		}
		customData = append(customData, customDatarecord)
	}
	return customData, nil
}

func (c *CsvIO) Write(writer io.Writer, customData []custom.CustomData) error {
	csvWriter := gocsv.DefaultCSVWriter(writer)
	defer csvWriter.Flush()
	var records []customIo.CustomIoData
	for _, record := range customData {
		records = append(records, c.FormatCustomDataToCustomIoData(record))
	}
	if err := gocsv.MarshalCSV(records, csvWriter); err != nil {
		return err
	}
	return nil
}
