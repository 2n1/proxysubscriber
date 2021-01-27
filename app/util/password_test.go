package util

import "testing"

const testPwd="2n1@tuta.io"
func TestPassword(t *testing.T) {
	pwd, err := Password(testPwd)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(pwd)
}

func TestVerifyPassword(t *testing.T) {
	hashedPwds:=[]string{
		"$2a$10$v1qva4YV9Z6HY4SdC0u9jOf.vzDsDjeXZl0T0hzE0vO0LuAChXcQS",
		"$2a$10$B/rnc3sutaH4cGXAU0sKv.sg6/3cATe1tV/1s1mNM9THRXyFF/8/K",
	}
	for idx,p:=range hashedPwds {
		if !VerifyPassword(testPwd,p) {
			t.Errorf("verify failed %d",idx)
		}
	}
	pwds:=[]string{
		testPwd,
		"foobar",
		"2n1",
	}
	for idx,p:=range hashedPwds {
		for _,pp:=range pwds {
			if !VerifyPassword(pp, p) {
				t.Errorf("verify failed %s of %d", pp,idx)
			}
		}
	}
}