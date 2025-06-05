package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("hello", func() {
	Description("シンプルな挨拶を返すサービスです。")

	Method("sayHello", func() {
		Payload(String, "挨拶する相手の名前")
		Result(String, "挨拶メッセージ")

		HTTP(func() {
			GET("/hello/{name}")
		})
	})
})
