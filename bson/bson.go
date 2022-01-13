package bson

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func E(key string, value interface{}) bson.E {
	return bson.E{Key: key, Value: value}
}
func A(properties []string) bson.A {
	props := bson.A{}
	for i := 0; i < len(properties); i++ {
		props = append(props, properties[i])
	}
	return props
}

func Regex(pattern, options string) primitive.Regex {
	return primitive.Regex{
		Pattern: pattern,
		Options: options,
	}
}

type D = bson.D
type M = bson.M
