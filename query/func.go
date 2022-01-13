package query

import (
	"github.com/hoainguyen33/Go/bson"

	mbson "go.mongodb.org/mongo-driver/bson"
)

type Filter = bson.D
type Sorter = bson.D

// find by id
func FindByID(id string) bson.D {
	return bson.D{bson.E("$match", bson.D{bson.E("_id", id)})}
}

// find by properties...
func Find(properties ...mbson.E) bson.D {
	props := bson.D{}
	for i := 0; i < len(properties); i++ {
		props = append(props, properties[i])
	}

	return bson.D{bson.E("$match", props)}
}

// find in array
func FindIn(filter map[string]interface{}) bson.D {
	props := bson.M{}
	for i, s := range filter {
		props[i] = bson.M{"$in": s}
	}
	return bson.D{bson.E("$match", props)}
}

func SearchRegex(key, pattern, options string) bson.D {
	return bson.D{bson.E("$match", bson.M{key: bson.Regex(pattern, options)})}
}

func Skip(offset int64) bson.D {
	return bson.D{
		bson.E("$skip", offset),
	}
}

func Limit(limit int64) bson.D {
	return bson.D{
		bson.E("$limit", limit),
	}
}

func Join(from, localField, foreignField string, as interface{}) bson.D {
	return bson.D{
		bson.E("$lookup", bson.D{
			bson.E("from", from),
			bson.E("localField", localField),
			bson.E("foreignField", foreignField),
			bson.E("as", as),
		}),
	}
}

func Unwind(path string) bson.D {
	return bson.D{
		bson.E("$unwind", path),
	}
}

func Group(properties ...mbson.E) bson.D {
	props := bson.D{}
	for i := 0; i < len(properties); i++ {
		props = append(props, properties[i])
	}

	return bson.D{bson.E("$group", props)}
}

func Joins(from, localField, foreignField string, let mbson.E, pipeline bson.D, as string) bson.D {
	return bson.D{
		bson.E("$lookup", bson.D{
			bson.E("from", from),
			bson.E("localField", localField),
			bson.E("foreignField", foreignField),
			bson.E("let", let),
			bson.E("foreignField", pipeline),
			bson.E("as", as),
		}),
	}
}

func Select(properties ...string) bson.D {
	props := bson.D{}
	for i := 0; i < len(properties); i++ {
		props = append(props, bson.E(properties[i], 1))
	}

	return bson.D{bson.E("$project", props)}
}

func Count(fieldName string) bson.D {
	return bson.D{bson.E("$count", fieldName)}
}

func Facet(outputFields bson.M) bson.D {
	return bson.D{bson.E("$facet", outputFields)}
}
