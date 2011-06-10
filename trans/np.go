package trans

import (
    dbi "github.com/thomaslee/go-dbi"
    _ "github.com/thomaslee/go-dbd-sqlite"
//    vector "container/vector"
    "fmt"
)

type BasicListing struct {
    ProductName string
    ModelNumber string
    List string
    Cost string
}

func (listing *BasicListing) String() string {
    s := fmt.Sprintf(`{ product_name: '%s',
  model_number: '%s',
  list: '%s',
  cost: '%s'}
	`, listing.ProductName, listing.ModelNumber, listing.List,
	listing.Cost)
    return s
}

func NpToAmazon () ([]map[string]string) {
    npconn, err := dbi.Connect("sqlite://./np.sqlite")
    if err != nil {
	fmt.Printf("error: connecting to np.sqlite: %s\n", err.String())
	return nil
    }
    defer npconn.Close()

    amaconn, err := dbi.Connect("sqlite://./products.sqlite")
    if err != nil {
	fmt.Printf("error: connecting to products.sqlite: %s\n",
	    err.String())
	return nil
    }
    defer amaconn.Close()

    rs, err := npconn.Query(`SELECT product_name, model_number, list,
	cost FROM items`)
    if err != nil {
	fmt.Printf("error: reading from np.sqlite: %s\n", err.String())
	return nil
    }
    defer rs.Close()

    ret := []map[string]string {}
    var cur []map[string]string
    for rs.Next() {
	var product_name string
	var model_number string
	var list string
	var cost string

	err = rs.Scan(&product_name, &model_number, &list, &cost)
	if err != nil {
	    fmt.Printf("error: %s\n", err.String())
	}

	cur = []map[string]string{ {
	    "product_name": product_name,
	    "model_number": model_number,
	    "list":	    list,
	    "cost":	    cost,
	} }

	newslice := make([]map[string]string, len(ret) + len(cur))
	copy(newslice, ret)
	copy(newslice[len(ret):], cur)
	ret = newslice

	/*item := &BasicListing{ ProductName: product_name, ModelNumber:
	    model_number, List: list, Cost: cost}

	vec.Push(item)*/

    }

//    fmt.Print(ret)
    return ret

/*
    for i := 0; i < vec.Len(); i++ {
	el := vec.At(i);
	fmt.Print(el,"\n");
    }
    */
}
