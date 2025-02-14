package config

import (
	"crypto"
	"crypto/ed25519"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sosodev/duration"
)

type jwtPrivateKey struct {
	Id  int8
	Key crypto.PrivateKey
}
type JWTConfig struct {
	PublicKeys map[int8]crypto.PublicKey
	PrivateKey jwtPrivateKey
	AccessTokenTTL time.Duration
	RefreshTokenTTL time.Duration
}
func createJWTPrivateKeyParser(publicKeys map[int8]crypto.PublicKey) parseFunc[jwtPrivateKey] {
	return func (value string) (*jwtPrivateKey,error){
		parts:= strings.Split(value,":")
		i64,error:= strconv.ParseInt(parts[0],10,8)
		if error != nil {
			return nil, errors.New(
				fmt.Sprintf("Failed to parsed id as int8 %s",error))
		}
		id := int8(i64)
		keyString :=  parts[1]
		key,error:= jwt.ParseEdPrivateKeyFromPEM([]byte(keyString))
		if error != nil {
			return nil, errors.New( fmt.Sprintf("Failed to parsed key %s",error))
		}
		message:= []byte("asdjdf")
		signature:= ed25519.Sign(key.(ed25519.PrivateKey),message)
		corresPondingPublicKey,ok := publicKeys[id]
		if !ok {
			panic(fmt.Sprintf("No corresponding public key for id %d",id))
	    }
		err := ed25519.VerifyWithOptions(
			corresPondingPublicKey.(ed25519.PublicKey),
			message,signature,&ed25519.Options{},
		)
		if err != nil {
			panic(fmt.Sprintf(
				"jwtPrivateKey: Corresponding private key cannot" + "verify test message %s",err,
			))
		}
		return &jwtPrivateKey{Id:id,Key:key},nil

	}
}

func parseJwtPublicKey (value string) (int8,*crypto.PublicKey,error){
	parts:= strings.Split(value,":")
	i64,err:= strconv.ParseInt(parts[0],10,8)
	if err != nil {
		return 0, nil,errors.New(fmt.Sprintf("Failed to parsed id as int8 %s",err))
	}
	id:= int8(i64)
	keyString:= parts[1]
	key, err := jwt.ParseEdPublicKeyFromPEM([]byte(keyString))
	if err != nil {
		return 0,nil,err
	}
	return id, &key, err
}
func parseJwtPublicKeys (value string) (*map[int8]crypto.PublicKey,error){
	parts:= strings.Split(value,";")
	keys:= make(map[int8]crypto.PublicKey)
    for index, parts:= range(parts){
		id,key,err:= parseJwtPublicKey(parts)
		if err != nil {
			return nil,errors.New(
				fmt.Sprintf("[%d] key is invalid: %e",index,err),
			)
		}
		keys[id] = *key
	}
	return &keys, nil
}
func parseTTL (value string) (*time.Duration, error){
	val,err:= duration.Parse(value)
	if err != nil {
		return nil, err
	}
	td:= val.ToTimeDuration()
	return &td,nil
}

func createJWTConfig() JWTConfig {
	publicKeys :=  getEnv("JWT_PUBLIC_KEYS",false,parseJwtPublicKeys)
	return JWTConfig{
		PublicKeys:publicKeys,
		PrivateKey: getEnv("JWT_PRIVATE_KEY",false,createJWTPrivateKeyParser(publicKeys)),
		AccessTokenTTL: getEnv("JWT_ACCESS_TOKEN_TTL",false,parseTTL),
		RefreshTokenTTL: getEnv("JWT_REFRESH_TOKEN_TTL",false,parseTTL),

	}
}
