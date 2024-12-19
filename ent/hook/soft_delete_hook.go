package hook

import (
	"context"
	"fmt"
	"reflect"

	"entgo.io/ent/dialect/sql"
)

// AddSoftDeleteFilter applies a query filter to exclude soft-deleted records.
func AddSoftDeleteFilter(client interface{}) error {
	clientVal := reflect.ValueOf(client)

	// Check if the client has a "Use" method
	useMethod := clientVal.MethodByName("Use")
	if !useMethod.IsValid() {
		return fmt.Errorf("client does not have a 'Use' method, type: %T", client)
	}

	fmt.Println("Found 'Use' method, applying middleware")

	// Define the middleware function dynamically
	middleware := reflect.MakeFunc(
		reflect.TypeOf(func(next interface{}) interface{} { return nil }),
		func(args []reflect.Value) (results []reflect.Value) {
			next := args[0]
			return []reflect.Value{
				reflect.MakeFunc(
					reflect.TypeOf(func(ctx context.Context, mutation interface{}) (interface{}, error) { return nil, nil }),
					func(innerArgs []reflect.Value) []reflect.Value {
						mutation := innerArgs[1]

						// Apply global filters to exclude soft-deleted records
						mutationVal := reflect.ValueOf(mutation)
						wherePMethod := mutationVal.MethodByName("WhereP")
						if wherePMethod.IsValid() {
							wherePMethod.Call([]reflect.Value{
								reflect.ValueOf(func(s *sql.Selector) {
									s.Where(sql.IsNull(s.C("deleted_at")))
								}),
							})
						}

						// Call the next mutator
						results := reflect.ValueOf(next).Call(innerArgs)
						if len(results) != 2 {
							return []reflect.Value{
								reflect.Zero(innerArgs[1].Type()), // Return nil for value
								reflect.ValueOf(fmt.Errorf("unexpected number of return values from next mutator")),
							}
						}

						return results
					},
				),
			}
		},
	)

	// Call the "Use" method with the dynamically created middleware
	useMethod.Call([]reflect.Value{middleware})

	fmt.Println("Soft-delete filter applied successfully")
	return nil
}
