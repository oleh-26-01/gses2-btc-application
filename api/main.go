package api

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

var subscriptionsPath = "./api/subscriptions.json"
var subscriptions = LoadEmails(subscriptionsPath)

func ConnectTo(app *iris.Application, route string) {
	app.Get(route+"/rate", func(ctx iris.Context) {
		price := RateResponse{GetBTCPrice()}
		_ = ctx.JSON(price)
	})

	app.Post(route+"/subscribe", func(ctx iris.Context) {
		var response EmailResponse

		err := ctx.ReadJSON(&response)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
		} else {
			isUnique := IsUnique(subscriptions, response)
			if isUnique {
				subscriptions = append(subscriptions, response)
				ctx.StatusCode(iris.StatusOK)
				_ = ctx.JSON(response)
				SaveEmails(subscriptionsPath, subscriptions)
			} else {
				ctx.StatusCode(iris.StatusConflict)
			}
		}
	})

	app.Post(route+"/sendEmails", func(ctx iris.Context) {
		result := SendEmails("")
		if result {
			fmt.Println("Emails were sent")
			ctx.StatusCode(iris.StatusOK)
		} else {
			fmt.Println("Error sending emails")
			ctx.StatusCode(iris.StatusInternalServerError)
		}
	})
}
