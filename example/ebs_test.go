package example

import (
	"context"
	"fmt"
	"log"
	"time"

	"git.supremind.info/products/atom/didiyunclient/pkg"
)

func ExampleCreateDelete_Ebs() {
	ctx := context.Background()
	token := "" // TODO: set your api token
	c, e := pkg.New(&pkg.Config{Token: token, Timeout: 5 * time.Second})
	if e != nil {
		log.Fatalln(e)
	}
	ebs := c.Ebs()
	id, e := ebs.Create(ctx, "gz", "gz02", "Example_Ebs", "SSD", 20) // name is not unique
	if e != nil {
		log.Fatalln(e)
	}

	if e = ebs.Delete(ctx, id); e != nil {
		log.Fatalln(e)
	}

	fmt.Println("Ebs created & deleted ok")
	// Output: Ebs created & deleted ok
}