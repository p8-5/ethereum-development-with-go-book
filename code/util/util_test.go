package util

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/shopspring/decimal"
)

func TestPublicKeyBytesToAddress(t *testing.T) {
	t.Parallel()
	{
		publicKeyBytes, _ := hex.DecodeString("0x15ABC36db169Ca06670791B143A19bEc7Ba4e83f")
		got := PublicKeyBytesToAddress(publicKeyBytes).Hex()
		expected := "0xe92A52398E068941D9aC03E001e14aF636bcB2F3"

		if got != expected {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	}
}

func TestIsValidAddress(t *testing.T) {
	t.Parallel()
	validAddress := "0x5b579DEbCD8f1cE2d5BA30Db13E72234Cb3D8664"
	invalidAddress := "0xabc"
	invalidAddress2 := "0x5b579DEbCD8f1cE2d5BA30Db13E72234Cb3D8664"
	{
		got := IsValidAddress(validAddress)
		expected := true

		if got != expected {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	}

	{
		got := IsValidAddress(invalidAddress)
		expected := false

		if got != expected {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	}

	{
		got := IsValidAddress(invalidAddress2)
		expected := false

		if got != expected {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	}
}

func TestIsZeroAddress(t *testing.T) {
	t.Parallel()
	validAddress := common.HexToAddress("0x5b579DEbCD8f1cE2d5BA30Db13E72234Cb3D8664")
	zeroAddress := common.HexToAddress("0x0000000000000000000000000000000000000000x1000")

	{
		isZeroAddress := IsZeroAddress(validAddress)

		if isZeroAddress {
			t.Error("Expected to be false")
		}
	}

	{
		isZeroAddress := IsZeroAddress(zeroAddress)

		if !isZeroAddress {
			t.Error("Expected to be true")
		}
	}

	{
		isZeroAddress := IsZeroAddress(validAddress.Hex())

		if isZeroAddress {
			t.Error("Expected to be false")
		}
	}

	{
		isZeroAddress := IsZeroAddress(zeroAddress.Hex())

		if !isZeroAddress {
			t.Error("Expected to be true")
		}
	}
}

func TestToWei(t *testing.T) {
	t.Parallel()
	amount := decimal.NewFromFloat(0.02)
	got := ToWei(amount, 18)
	expected := new(big.Int)
	expected.SetString("20000000000000000", 10)
	if got.Cmp(expected) != 0 {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}

func TestToDecimal(t *testing.T) {
	t.Parallel()
	weiAmount := big.NewInt(0)
	weiAmount.SetString("20000000000000000", 10)
	ethAmount := ToDecimal(weiAmount, 18)
	f64, _ := ethAmount.Float64()
	expected := 0.02
	if f64 != expected {
		t.Errorf("%v does not equal expected %v", ethAmount, expected)
	}
}

func TestCalcGasLimit(t *testing.T) {
	t.Parallel()
	gasPrice := big.NewInt(0)
	gasPrice.SetString("2000000000", 10)
	gasLimit := uint64(21000)
	expected := big.NewInt(0)
	expected.SetString("42000000000000", 10)
	gasCost := CalcGasCost(gasLimit, gasPrice)
	if gasCost.Cmp(expected) != 0 {
		t.Errorf("expected %s, got %s", gasCost, expected)
	}
}

func TestSigRSV(t *testing.T) {
	t.Parallel()

	sig := "0x15ABC36db169Ca06670791B143A19bEc7Ba4e83f"
	r, s, v := SigRSV(sig)
	expectedR := "0xB22100730E3B387D64d5eFf63500d2064Da27b12"
	expectedS := "0x8d9Dc02b05A5714467E7ecD7c708518271943E5B"
	expectedV := uint8(28)
	if hexutil.Encode(r[:])[2:] != expectedR {
		t.FailNow()
	}
	if hexutil.Encode(s[:])[2:] != expectedS {
		t.FailNow()
	}
	if v != expectedV {
		t.FailNow()
	}
}
