package csv

import (
	"bytes"
	"strings"
	"testing"

	"github.com/kis9a/cryptact/pkg/custom"
	"github.com/stretchr/testify/assert"
)

var csvData = `timestamp,action,source,base,volume,price,counter,fee,feeccy,comment
2017/01/02 12:23:00,BUY,any,ETH,2,0.5,BTC,0.01,ETH,2.1.1.Buy and Sell
2017/01/02 12:23:00,BUY,any,BCH,10,0,JPY,0,JPY,2.1.2.Hard folk
`

func TestCsvIO_Read(t *testing.T) {
	reader := strings.NewReader(csvData)
	csvIO := NewCsvIO()
	customData, err := csvIO.Read(reader)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(customData))

	data1 := customData[0]
	assert.Equal(t, "2017/01/02 12:23:00", data1.Timestamp.Format(csvIO.DateFormat))
	assert.Equal(t, custom.Action("BUY"), data1.Action)
	assert.Equal(t, "any", data1.Source)
	assert.Equal(t, "ETH", data1.Base)
	actual, _ := data1.Price.Float64()
	assert.Equal(t, float64(0.5), actual)
}

func TestCsvIO_Write(t *testing.T) {
	reader := strings.NewReader(csvData)
	csvIO := NewCsvIO()
	customData, _ := csvIO.Read(reader)

	var fileBuf bytes.Buffer
	err := csvIO.Write(&fileBuf, customData)
	assert.NoError(t, err)
	assert.NoError(t, err)
	assert.Equal(t, csvData, fileBuf.String())
}

// func TestCsvIO_Write(t *testing.T) {
// 	reader := strings.NewReader(csvData)
// 	csvIO := NewCsvIO()
// 	customData, _ := csvIO.Read(reader)

// 	var fileBuf bytes.Buffer
// 	err := csvIO.Write(&fileBuf, customData)
// 	assert.NoError(t, err)
// 	assert.NoError(t, err)
// 	assert.Equal(t, csvData, fileBuf.String())
// }
