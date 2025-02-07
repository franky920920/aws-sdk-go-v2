// Code generated by smithy-go-codegen DO NOT EDIT.

//go:build snapshot

package keyspaces

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/smithy-go/middleware"
	"io"
	"io/fs"
	"os"
	"testing"
)

const ssprefix = "snapshot"

type snapshotOK struct{}

func (snapshotOK) Error() string { return "error: success" }

func createp(path string) (*os.File, error) {
	if err := os.Mkdir(ssprefix, 0700); err != nil && !errors.Is(err, fs.ErrExist) {
		return nil, err
	}
	return os.Create(path)
}

func sspath(op string) string {
	return fmt.Sprintf("%s/api_op_%s.go.snap", ssprefix, op)
}

func updateSnapshot(stack *middleware.Stack, operation string) error {
	f, err := createp(sspath(operation))
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write([]byte(stack.String())); err != nil {
		return err
	}
	return snapshotOK{}
}

func testSnapshot(stack *middleware.Stack, operation string) error {
	f, err := os.Open(sspath(operation))
	if errors.Is(err, fs.ErrNotExist) {
		return snapshotOK{}
	}
	if err != nil {
		return err
	}
	defer f.Close()
	expected, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	if actual := stack.String(); actual != string(expected) {
		return fmt.Errorf("%s != %s", expected, actual)
	}
	return snapshotOK{}
}
func TestCheckSnapshot_CreateKeyspace(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreateKeyspace(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "CreateKeyspace")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_CreateTable(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreateTable(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "CreateTable")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_CreateType(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreateType(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "CreateType")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_DeleteKeyspace(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteKeyspace(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "DeleteKeyspace")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_DeleteTable(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteTable(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "DeleteTable")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_DeleteType(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteType(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "DeleteType")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_GetKeyspace(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetKeyspace(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "GetKeyspace")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_GetTable(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetTable(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "GetTable")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_GetTableAutoScalingSettings(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetTableAutoScalingSettings(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "GetTableAutoScalingSettings")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_GetType(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetType(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "GetType")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListKeyspaces(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListKeyspaces(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListKeyspaces")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListTables(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListTables(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListTables")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListTagsForResource(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListTagsForResource(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListTagsForResource")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_ListTypes(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListTypes(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "ListTypes")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_RestoreTable(t *testing.T) {
	svc := New(Options{})
	_, err := svc.RestoreTable(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "RestoreTable")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_TagResource(t *testing.T) {
	svc := New(Options{})
	_, err := svc.TagResource(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "TagResource")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_UntagResource(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UntagResource(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "UntagResource")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestCheckSnapshot_UpdateTable(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateTable(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return testSnapshot(stack, "UpdateTable")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}
func TestUpdateSnapshot_CreateKeyspace(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreateKeyspace(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "CreateKeyspace")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_CreateTable(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreateTable(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "CreateTable")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_CreateType(t *testing.T) {
	svc := New(Options{})
	_, err := svc.CreateType(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "CreateType")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_DeleteKeyspace(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteKeyspace(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "DeleteKeyspace")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_DeleteTable(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteTable(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "DeleteTable")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_DeleteType(t *testing.T) {
	svc := New(Options{})
	_, err := svc.DeleteType(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "DeleteType")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_GetKeyspace(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetKeyspace(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "GetKeyspace")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_GetTable(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetTable(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "GetTable")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_GetTableAutoScalingSettings(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetTableAutoScalingSettings(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "GetTableAutoScalingSettings")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_GetType(t *testing.T) {
	svc := New(Options{})
	_, err := svc.GetType(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "GetType")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListKeyspaces(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListKeyspaces(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListKeyspaces")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListTables(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListTables(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListTables")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListTagsForResource(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListTagsForResource(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListTagsForResource")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_ListTypes(t *testing.T) {
	svc := New(Options{})
	_, err := svc.ListTypes(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "ListTypes")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_RestoreTable(t *testing.T) {
	svc := New(Options{})
	_, err := svc.RestoreTable(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "RestoreTable")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_TagResource(t *testing.T) {
	svc := New(Options{})
	_, err := svc.TagResource(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "TagResource")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_UntagResource(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UntagResource(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "UntagResource")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}

func TestUpdateSnapshot_UpdateTable(t *testing.T) {
	svc := New(Options{})
	_, err := svc.UpdateTable(context.Background(), nil, func(o *Options) {
		o.APIOptions = append(o.APIOptions, func(stack *middleware.Stack) error {
			return updateSnapshot(stack, "UpdateTable")
		})
	})
	if _, ok := err.(snapshotOK); !ok && err != nil {
		t.Fatal(err)
	}
}
