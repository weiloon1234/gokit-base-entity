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
		return fmt.Errorf("client does not have a 'Use' method")
	}

	// Define the middleware function to exclude soft-deleted records
	middleware := func(next interface{}) interface{} {
		return func(ctx context.Context, mutation interface{}) (interface{}, error) {
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
			results := reflect.ValueOf(next).Call([]reflect.Value{reflect.ValueOf(ctx), mutationVal})
			if len(results) != 2 {
				return nil, fmt.Errorf("unexpected number of return values from next mutator")
			}

			// Return the value and error
			return results[0].Interface(), results[1].Interface().(error)
		}
	}

	// Call the "Use" method with the middleware
	useMethod.Call([]reflect.Value{reflect.ValueOf(middleware)})

	return nil
}
