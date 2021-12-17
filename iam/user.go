package iam

import (
	"crypto/md5"
	"icesos/kv"
)

type User string

func (user User) UserIAMKey() string {
	return string(user) + userIAMKey
}

func (user User) SetIAMKey() string {
	return string(user) + setIAMKey
}

func (user User) IsOwnSet(set Set) bool {
	setIAM := SetIAM{
		User: user,
		Set:  set,
	}

	member, err := setIAM.EncodeProto()
	if err != nil {
		return false
	}

	ret, _ := kv.Client.SIsMember(setIAM.Key(), member)
	return ret
}

func (user User) AddSet(set Set) error {
	setIAM := SetIAM{
		User: user,
		Set:  set,
	}

	member, err := setIAM.EncodeProto()
	if err != nil {
		return err
	}

	return kv.Client.SAdd(setIAM.Key(), member)
}

func (user User) Identify(sk string) bool {
	val, err := kv.Client.KvGet(user.UserIAMKey())
	if err != nil {
		return false
	}

	userIAM, err := DecodeUserIAMProto(val)
	if err != nil {
		return false
	}

	return userIAM.SecretKey == md5.Sum([]byte(sk))
}

func (user User) CreateUser(sk string) error {
	userIAM := UserIAM{
		User:      user,
		SecretKey: md5.Sum([]byte(sk)),
	}

	b, err := userIAM.EncodeProto()
	if err != nil {
		return err
	}

	return kv.Client.KvPut(userIAM.Key(), b)
}

func (user User) DeleteUser() (bool, error) {
	ret, err := kv.Client.KvDelete(user.UserIAMKey())
	if err != nil || ret == false {
		return false, err
	}

	_, err = kv.Client.SDelete(user.SetIAMKey())
	return true, err
}