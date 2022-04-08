package mathutil

import (
	math2 "github.com/ethereum/go-ethereum/common/math"
	"math"
	"math/big"
	"strconv"
	"strings"
)

var Wei18 =big.NewInt(int64(math.Pow10(18)))

func FloatToString(f float64) string {

	fStr := strconv.FormatFloat(f, 'f', -1, 64)

	return fStr

}
func Int64ToString(i int64) string {
	return strconv.FormatInt(i,10)
}
func StringToInt64(s string) (int64 ,error) {
	return strconv.ParseInt(s,10,64)
}

func StringToFloat(f string) float64 {

	n, err := strconv.ParseFloat(f, 64)

	if err != nil {
		return 0
	}
	return n

}

func FloatToWei(f float64) *big.Int {

	weiCount := 18

	beth := math2.BigPow(10, int64(weiCount))

	rs := new(big.Float).Mul(big.NewFloat(f), new(big.Float).SetInt(beth))

	w:=big.NewInt(0)

	rs.Int(w)

	return w

}
func WeiToFloat(w *big.Int) float64 {
	weiCount := 18

	beth := math2.BigPow(10, int64(weiCount))

	re :=new(big.Rat).Quo(new(big.Rat).SetInt(w),new(big.Rat).SetInt(beth))

	fe,_ := re.Float64()


	return fe

}

func GetFloatDecimal(f float64) int {

	fStr := strconv.FormatFloat(f, 'f', -1, 64)

	_, endCount := GetFloatStrDecimal(fStr)

	return endCount

}
func ParseFloat(f float64,n int) float64 {

	fStr := strconv.FormatFloat(f, 'f', n, 64)


	return StringToFloat(fStr)

}

func GetFloatStrDecimal(f_str string) (headCount int, endCount int) {

	strLen := len(f_str)

	pIndex := strings.Index(f_str, ".")

	if pIndex >= 0 {

		return pIndex, strLen - pIndex

	}

	return strLen, 0
}
