package main

import (
	"context"
	"fmt"
)

func main() {

	autenticar(context.Background())

}

func autenticar(ctx context.Context) {

	ctx = context.WithValue(ctx, "env", "sandbox")

	func1(ctx)

}

func func1(ctx context.Context) {

	fmt.Println("hola desde func1")

	func2(ctx)

}

func func2(ctx context.Context) {

	fmt.Println("hola desde func2")

	env := "live"
	if v := ctx.Value("env"); v != nil {
		envString, ok := v.(string)
		if ok {
			env = envString
		}
	}

	fmt.Println("env: ", env)

}
