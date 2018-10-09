package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"reflect"
)

type Resource struct {
	Title string `json:"title"`
}

func main() {
	var r Resource
	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}
	query := elastic.NewMatchQuery("title", "简称")
	result, _ := client.Search("darion_index").Query(query).Do(context.Background())

	for _, item := range result.Each(reflect.TypeOf(r)) {
		if resource, ok := item.(Resource); ok {
			fmt.Println(resource.Title)
		}
	}
	fmt.Println("--------------------------------------------------------------------------------------------------------------------")
	multiQuery := elastic.NewMultiMatchQuery(
		"简称",
		"title",
	).Type("most_fields")

	multiResult, _ := client.Search("darion_index").Query(multiQuery).Do(context.Background())
	//for _, item := range multiResult.Each(reflect.TypeOf(r)) {
	//	if resource, ok := item.(Resource); ok {
	//		fmt.Printf("%d %f %s\n", resource.ID, resource.Score, resource.Title)
	//	}
	//}

	for _, hit := range multiResult.Hits.Hits {
		err := json.Unmarshal(*hit.Source, &r)
		if err != nil {
			panic(0)
		}

		fmt.Printf("%s %f %s \n", hit.Id, *hit.Score, r.Title)
	}
	fmt.Println(multiResult.Hits.TotalHits)
}
