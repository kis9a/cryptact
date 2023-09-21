package custom

import (
	"math/big"
	"testing"
	"time"
)

func TestCustom(t *testing.T) {
	t.Run("TestCustomParsing", func(t *testing.T) {
		timestamp := time.Now()
		action := ActionBuy
		source := "Exchange"
		base := "ANY"
		volume := big.NewFloat(1.0)
		price := big.NewFloat(50000.0)
		counter := "USD"
		fee := big.NewFloat(0.5)
		feeCcy := "USD"
		comment := "Sample comment"

		customData := CustomData{
			Timestamp: timestamp,
			Action:    action,
			Source:    source,
			Base:      base,
			Volume:    volume,
			Price:     price,
			Counter:   counter,
			Fee:       fee,
			FeeCcy:    feeCcy,
			Comment:   comment,
		}

		if customData.Timestamp != timestamp ||
			customData.Action != action ||
			customData.Source != source ||
			customData.Base != base ||
			customData.Volume.Cmp(volume) != 0 ||
			customData.Price.Cmp(price) != 0 ||
			customData.Counter != counter ||
			customData.Fee.Cmp(fee) != 0 ||
			customData.FeeCcy != feeCcy ||
			customData.Comment != comment {
			t.Errorf("Custom parsing failed")
		}
	})

	t.Run("TestActionStringConversion", func(t *testing.T) {
		action := ActionBuy
		actionStr := action.StringValue()

		if actionStr != "BUY" {
			t.Errorf("Action string conversion failed")
		}
	})
}
