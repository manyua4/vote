package model

import (
	"context"
	"fmt"
	"testing"
)

func TestGetVoteV1(t *testing.T) {
	NewMysql()
	ret := GetVoteV1(1)
	fmt.Printf("ret:%+v\n", ret)
}
func TestGetVoteV2(t *testing.T) {
	NewMysql()
	ret, _ := GetVoteV2(1)
	fmt.Printf("ret:%+v\n", ret)
}

func TestGetVoteV3(t *testing.T) {
	NewMysql()
	ret, _ := GetVoteV3(1)
	fmt.Printf("ret:%+v\n", ret)
}
func TestGetVoteV4(t *testing.T) {
	NewMysql()
	ret, _ := GetVoteV4(1)
	fmt.Printf("ret:%+v\n", ret)
}
func TestGetVoteV5(t *testing.T) {
	NewMysql()
	ret, _ := GetVoteV5(1)
	fmt.Printf("ret:%+v\n", ret)
}

func TestGetVoteHistoryV1(t *testing.T) {
	NewMysql()
	NewRdb()
	GetVoteHistoryV1(context.TODO(), 2, 1)
}
func TestGetJwt(t *testing.T) {
	str, _ := GetJwt(1, "香香编程")
	fmt.Printf("str:%s", str)
}

func TestCheckJwt(t *testing.T) {
	str := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiTmFtZSI6Iummmemmmee8lueoiyIsImlzcyI6Iummmemmmee8lueoiyIsInN1YiI6IuWQjuWLpOmDqOenpuW4iOWChSIsImF1ZCI6WyJBbmRyb2lkIiwiSU9TIiwiSDUiXSwiZXhwIjoxNzAzNTkwNTg4LCJuYmYiOjE3MDM1ODY5OTgsImlhdCI6MTcwMzU4Njk4OCwianRpIjoiVGVzdC0xIn0.9WaYHAr_2cyCnzoICMyx5vbDWFJxydhE8NhCh4Eye60"
	token, _ := CheckJwt(str)
	fmt.Printf("token:%+v\n", token)
}
