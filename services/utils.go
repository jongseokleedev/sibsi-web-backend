package services

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// global variable
var dbName = "sibsiDB"

func GetCollection(client *mongo.Client, colName string) *mongo.Collection {
	return client.Database(dbName).Collection(colName)
}

//func InsertDocuments(client *mongo.Client, ctx context.Context, docs []interface{}) {
//	collection := client.Database("sibsiDB").Collection("provider")
//
//	insertManyResult, err := collection.InsertMany(ctx, docs)
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Printf("Inserted documents with IDs: %v", insertManyResult.InsertedIDs)
//}
//
//func CreateAnimals() []interface{} {
//	encodedImages := EncodeFileToBase64()
//	return []interface{}{
//		bson.D{{"index", 1}, {"OriginalName", "badger"}, {"LeftPartName", "bad"}, {"RightPartName", "ger"}, {"BaseEncodedImage", encodedImages["badger"]}},
//		bson.D{{"index", 2}, {"OriginalName", "bear"}, {"LeftPartName", "be"}, {"RightPartName", "ar"}, {"BaseEncodedImage", encodedImages["bear"]}},
//		bson.D{{"index", 3}, {"OriginalName", "cat"}, {"LeftPartName", "ca"}, {"RightPartName", "t"}, {"BaseEncodedImage", encodedImages["cat"]}},
//		bson.D{{"index", 4}, {"OriginalName", "cheetah"}, {"LeftPartName", "chee"}, {"RightPartName", "tah"}, {"BaseEncodedImage", encodedImages["cheetah"]}},
//		bson.D{{"index", 5}, {"OriginalName", "chicken"}, {"LeftPartName", "chi"}, {"RightPartName", "ken"}, {"BaseEncodedImage", encodedImages["chicken"]}},
//		bson.D{{"index", 6}, {"OriginalName", "chinchilla"}, {"LeftPartName", "chin"}, {"RightPartName", "chilla"}, {"BaseEncodedImage", encodedImages["chinchilla"]}},
//		bson.D{{"index", 7}, {"OriginalName", "clownfish"}, {"LeftPartName", "clown"}, {"RightPartName", "fish"}, {"BaseEncodedImage", encodedImages["clownfish"]}},
//		bson.D{{"index", 8}, {"OriginalName", "crane"}, {"LeftPartName", "cra"}, {"RightPartName", "ne"}, {"BaseEncodedImage", encodedImages["crane"]}},
//		bson.D{{"index", 9}, {"OriginalName", "duck"}, {"LeftPartName", "du"}, {"RightPartName", "ck"}, {"BaseEncodedImage", encodedImages["duck"]}},
//		bson.D{{"index", 10}, {"OriginalName", "frog"}, {"LeftPartName", "fr"}, {"RightPartName", "og"}, {"BaseEncodedImage", encodedImages["frog"]}},
//		bson.D{{"index", 11}, {"OriginalName", "hedgehog"}, {"LeftPartName", "hedge"}, {"RightPartName", "hog"}, {"BaseEncodedImage", encodedImages["hedgehog"]}},
//		bson.D{{"index", 12}, {"OriginalName", "hippo"}, {"LeftPartName", "hi"}, {"RightPartName", "ppo"}, {"BaseEncodedImage", encodedImages["hippo"]}},
//		bson.D{{"index", 13}, {"OriginalName", "horse"}, {"LeftPartName", "hor"}, {"RightPartName", "se"}, {"BaseEncodedImage", encodedImages["horse"]}},
//		bson.D{{"index", 14}, {"OriginalName", "kangaroo"}, {"LeftPartName", "kanga"}, {"RightPartName", "roo"}, {"BaseEncodedImage", encodedImages["kangaroo"]}},
//		bson.D{{"index", 15}, {"OriginalName", "koala"}, {"LeftPartName", "ko"}, {"RightPartName", "ala"}, {"BaseEncodedImage", encodedImages["koala"]}},
//		bson.D{{"index", 16}, {"OriginalName", "lion"}, {"LeftPartName", "li"}, {"RightPartName", "on"}, {"BaseEncodedImage", encodedImages["lion"]}},
//		bson.D{{"index", 17}, {"OriginalName", "monkey"}, {"LeftPartName", "mon"}, {"RightPartName", "key"}, {"BaseEncodedImage", encodedImages["monkey"]}},
//		bson.D{{"index", 18}, {"OriginalName", "mouse"}, {"LeftPartName", "mou"}, {"RightPartName", "se"}, {"BaseEncodedImage", encodedImages["mouse"]}},
//		bson.D{{"index", 19}, {"OriginalName", "owl"}, {"LeftPartName", "ow"}, {"RightPartName", "l"}, {"BaseEncodedImage", encodedImages["owl"]}},
//		bson.D{{"index", 20}, {"OriginalName", "pig"}, {"LeftPartName", "pi"}, {"RightPartName", "g"}, {"BaseEncodedImage", encodedImages["pig"]}},
//		bson.D{{"index", 21}, {"OriginalName", "pigeon"}, {"LeftPartName", "pi"}, {"RightPartName", "geon"}, {"BaseEncodedImage", encodedImages["pigeon"]}},
//		bson.D{{"index", 22}, {"OriginalName", "polarbear"}, {"LeftPartName", "polar"}, {"RightPartName", "bear"}, {"BaseEncodedImage", encodedImages["polarbear"]}},
//		bson.D{{"index", 23}, {"OriginalName", "rabbit"}, {"LeftPartName", "rab"}, {"RightPartName", "bit"}, {"BaseEncodedImage", encodedImages["rabbit"]}},
//		bson.D{{"index", 24}, {"OriginalName", "rat"}, {"LeftPartName", "ra"}, {"RightPartName", "t"}, {"BaseEncodedImage", encodedImages["rat"]}},
//		bson.D{{"index", 25}, {"OriginalName", "seaturtle"}, {"LeftPartName", "sea"}, {"RightPartName", "turtle"}, {"BaseEncodedImage", encodedImages["seaturtle"]}},
//		bson.D{{"index", 26}, {"OriginalName", "sheep"}, {"LeftPartName", "sh"}, {"RightPartName", "eep"}, {"BaseEncodedImage", encodedImages["sheep"]}},
//		bson.D{{"index", 27}, {"OriginalName", "swan"}, {"LeftPartName", "sw"}, {"RightPartName", "an"}, {"BaseEncodedImage", encodedImages["swan"]}},
//		bson.D{{"index", 28}, {"OriginalName", "trout"}, {"LeftPartName", "tro"}, {"RightPartName", "ut"}, {"BaseEncodedImage", encodedImages["trout"]}},
//		bson.D{{"index", 29}, {"OriginalName", "turtle"}, {"LeftPartName", "tur"}, {"RightPartName", "tle"}, {"BaseEncodedImage", encodedImages["turtle"]}},
//		bson.D{{"index", 30}, {"OriginalName", "zebra"}, {"LeftPartName", "ze"}, {"RightPartName", "bra"}, {"BaseEncodedImage", encodedImages["zebra"]}},
//	}
//}
//
//func EncodeFileToBase64() map[string]string {
//	pngFiles, err := filepath.Glob("./png/set_1/*.png")
//	if err != nil {
//		log.Fatal(err)
//	}
//	base64EncodedImages := make(map[string]string)
//	// Read each file and encode it to base64
//	for _, filename := range pngFiles {
//		data, err := ioutil.ReadFile(filename)
//		if err != nil {
//			log.Fatal(err)
//		}
//		encoded := base64.StdEncoding.EncodeToString(data)
//		basename := filepath.Base(filename)
//		basenameNoExt := strings.TrimSuffix(basename, ".png")
//		base64EncodedImages[basenameNoExt] = encoded
//	}
//
//	//fmt.Printf("%v", base64EncodedImages["chicken"])
//	return base64EncodedImages
//	// Now base64EncodedImages array contains base64 encoding of all your PNG files
//	// You can use it to insert into your BSON document
//
//}
