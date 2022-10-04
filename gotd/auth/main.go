// Binary auth implements authentication example for user using terminal.
package main

// import (
// 	"bufio"
// 	"context"
// 	"errors"
// 	"flag"
// 	"fmt"
// 	"os"
// 	"strings"

// 	"github.com/gotd/td/telegram"
// 	"github.com/gotd/td/telegram/auth"
// 	"github.com/gotd/td/tg"
// 	"go.uber.org/zap"
// 	"golang.org/x/term"
// )

// // noSignUp can be embedded to prevent signing up.
// type noSignUp struct{}

// func (c noSignUp) SignUp(ctx context.Context) (auth.UserInfo, error) {
// 	return auth.UserInfo{}, errors.New("not implemented")
// }

// func (c noSignUp) AcceptTermsOfService(ctx context.Context, tos tg.HelpTermsOfService) error {
// 	return &auth.SignUpRequired{TermsOfService: tos}
// }

// // termAuth implements authentication via terminal.
// type termAuth struct {
// 	noSignUp

// 	phone string
// }

// func (a termAuth) Phone(_ context.Context) (string, error) {
// 	return a.phone, nil
// }

// func (a termAuth) Password(_ context.Context) (string, error) {
// 	fmt.Print("Enter 2FA password: ")
// 	bytePwd, err := term.ReadPassword(0)
// 	if err != nil {
// 		return "", err
// 	}
// 	return strings.TrimSpace(string(bytePwd)), nil
// }

// func (a termAuth) Code(_ context.Context, _ *tg.AuthSentCode) (string, error) {
// 	fmt.Print("Enter code: ")
// 	code, err := bufio.NewReader(os.Stdin).ReadString('\n')
// 	if err != nil {
// 		return "", err
// 	}
// 	return strings.TrimSpace(code), nil
// }

// func main() {
// 	// context
// 	ctx := context.Background()
// 	// log
// 	log, _ := zap.NewDevelopment()

// 	// phoneNumber
// 	phoneNumber := "+6289635552242"

// 	// telegram client
// 	client, err := telegram.ClientFromEnvironment(telegram.Options{
// 		Logger: log,
// 	})
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	phone := flag.String("phone", phoneNumber, "phone number to authenticate")
// 	flag.Parse()

// 	// Setting up authentication flow helper based on terminal auth.
// 	flow := auth.NewFlow(
// 		termAuth{phone: *phone},
// 		auth.SendCodeOptions{},
// 	)

// 	client.Run(ctx, func(ctx context.Context) error {
// 		if err := client.Auth().IfNecessary(ctx, flow); err != nil {
// 			return err
// 		}

// 		log.Info("Success")

// 		return nil
// 	})
// }
