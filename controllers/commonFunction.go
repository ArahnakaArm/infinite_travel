package controllers

import (
	"fmt"
	"math/rand"
	"net/mail"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func getNumber12digit() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var codes [12]byte
	for i := 0; i < 12; i++ {
		codes[i] = uint8(48 + r.Intn(9))
	}
	return string(codes[:])
}

func getTicketNumber() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var codes [6]byte
	for i := 0; i < 6; i++ {
		codes[i] = uint8(48 + r.Intn(10))
	}

	return string(codes[:])
}

func getSeatNumber() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var codes [2]byte
	for i := 0; i < 2; i++ {
		codes[i] = uint8(48 + r.Intn(5))
	}

	return string(codes[:])
}

func generateTicket(flightName string) string {

	return fmt.Sprintf("%s-%s", flightName, getTicketNumber())
}

func generateSeat() string {

	randomSeatChar := 'A' + rune(rand.Intn(6))

	return fmt.Sprintf("%s-%s", string(randomSeatChar), getSeatNumber())
}

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func hashPassword(pass string) (hashedPassword string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err
}

func padNumberWithZero(value int) string {
	return fmt.Sprintf("%02d", value)
}
