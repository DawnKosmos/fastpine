package ftx

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/dawnkosmos/fastpine/exchange"
)

func TestSubscription(t *testing.T) {

	key := "aItoEjHCABXEuxg-C937wiOMZYMC2m0pr2pXYfQ3"
	secret := "v5wFE6jF599kA9y_agx9JegQqA_Xp_E45AFljlCZ"
	client := New(key, secret, "", []string{""})
	ch, err := client.OHCLV("BTC-PERP", 21600, exchange.DateToTime("01", "01", "2020"), time.Now().Unix())
	f, err := os.Create("test.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	var a int64 = 0
	for _, c := range ch {
		if c.StartTime.Unix() == a {
			fmt.Println("SCHEIße")
		}
		a = c.StartTime.Unix()
		_, err := f.WriteString(strconv.FormatInt(c.StartTime.Unix(), 10) + " " + strconv.Itoa(int(c.Open)) + " " + strconv.Itoa(int(c.Close)) + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println(ch[0].StartTime, len(ch))
	fmt.Println(err)
}
