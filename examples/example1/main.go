package main

import (
	"fmt"

	"github.com/joaogabriel01/gochestra"
	"github.com/joaogabriel01/gochestra/protocols"
)

type User struct {
	ID      string
	Details string
}

func main() {
	var err error
	units := map[string]protocols.StorageUnit[string, string]{}
	orchestrator := gochestra.NewOrchestrator[string, string](units, []string{})

	redisUnit := NewRedisStorageUnit()
	pgUnit := NewPostgresStorageUnit()

	err = orchestrator.AddUnit("redis", redisUnit)
	if err != nil {
		panic(err)
	}

	err = orchestrator.AddUnit("postgres", pgUnit)
	if err != nil {
		panic(err)
	}

	err = orchestrator.SetStandardOrder("redis", "postgres")
	if err != nil {
		panic(err)
	}
	userJson := `{"id": "1", "details": "details"}`

	saved, err := orchestrator.Save("1", userJson)
	if err != nil {
		panic(err)
	}
	// saved = ["redis", "postgres"]

	user, err := orchestrator.Get("1")
	if err != nil {
		panic(err)
	}
	// user = `{"id": "1", "details": "details"}`
	fmt.Println("after save: ", user)
	fmt.Println("saved: ", saved)

	err = orchestrator.Delete("1")
	if err != nil {
		panic(err)
	}

	user, err = orchestrator.Get("1")
	if err != nil {
		panic(err)
	}
	// user = ""
	fmt.Println("after delete: ", user)

}
