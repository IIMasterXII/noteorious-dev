package FOLDERS

import(
	"fmt"
	"testing"
	"google.golang.org/appengine/aetest"
)

func TestFolder(t *testing.T){
	ctx, done, err := aetest.NewContext()
	defer done()
	if err != nil {
		fmt.Println("PANIC in TestFolder")
		panic(1)
	}
	// No assertion to check, if it compiles it fulfills the interface.
	(&Folder{}).Key(ctx,"example_key")
}