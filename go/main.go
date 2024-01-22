package main

import (
	"fmt"

	"github.com/elliotchance/phpserialize"
)

func main() {
	out, err := phpserialize.Marshal("testing bigbox bigenvelope", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
	fmt.Println(string(out))

	// var in float32
	// err = phpserialize.Unmarshal(out, &in)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(in)

	// str := "a:6:{s:7:\"adapter\";s:9:\"Pdo_mysql\";s:4:\"host\";s:13:\"172.17.62.150\";s:4:\"port\";i:3306;s:6:\"dbname\";s:7:\"ds_demo\";s:8:\"username\";s:4:\"demo\";s:8:\"password\";s:280:\"cTdSZGJVVkFrODZtVURiQ0w3T3BCRVNzMDhYQW9lVmsxa09pN0ZMdXBldTNaby9zRmkrWnVBaXY0RnNFb1dwRlhvcTNPa1hLblFhT0V1cW8vOFF3bFpxR1E2KzhoMmg2RnhYUlQ4dGl4NDNmZlBsa2o5NjB5K3hLbjVNS0xBalNnN3duRDMyUVJNbmJLVm5xa0t5bUtwbnBXL21qRFI0VU5PcWNCbU1RV1FPOS85a0w1KzdIQVBoSStIQWpLYWh0OjrJrHZ0rNyNva9+X3Rk0yAF\";}"
	// str := "s:3:\"oke\";"
	// fmt.Println([]byte(str))
	// var cek string
	// err = phpserialize.Unmarshal([]byte(str), &cek)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(cek)

	// str, _ := phpserialize.Marshal("oke", nil)
	// fmt.Println(string(str))

}
