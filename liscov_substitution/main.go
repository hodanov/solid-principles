package main

import (
	"fmt"

	"main.go/service"
)

/*
乗り物（vehicle）
vehicleで「このスピードでここまで行く」という関数を定義したとする。
しかしvehicleは自動車、自転車、新幹線、飛行機、船、宇宙船などさまざまなものがある。
各vehicleはそれぞれ目的と手段が異なるため、同じクラスで定義し、同じ関数に依存するとおかしなことになる。
例えば下記文字列を出力するメソッドがvehicleに実装されていたとする。
```
Aは時速B[km/h]で移動する。
```
車やロードバイクは問題ないが、SpaceXのロケットは第一宇宙速度を突破するため秒速7.9[km/s]で打ち出されないといけない。なので、下記のような文言に修正したい。
```
spaceXのファルコンは秒速7.9[km/s]で移動する。
```
しかしvehicleは他からも呼び出されているので、直接変更するとそちらにも影響してしまう。
それを回避するため、interfaceを用いてメソッドを分割する。
*/

func main() {
	var car service.VehicleInterface = &service.Car{
		Name:  "ランドクルーザー",
		Speed: 60,
	}
	var bicycle service.VehicleInterface = &service.Bicycle{
		Name:  "ロードバイク",
		Speed: 10,
	}
	var spaceX service.VehicleInterface = &service.SpaceX{
		Name:  "ファルコン",
		Speed: 7.9,
	}

	fmt.Println(car.Ride() + car.Transport("新潟"))
	fmt.Println(bicycle.Ride() + bicycle.Transport("近所のコンビニ"))
	fmt.Println(spaceX.Ride() + spaceX.Transport("宇宙"))
}
