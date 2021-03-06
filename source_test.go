package locerr

import (
	"strings"
	"testing"
)

func TestReadFromFile(t *testing.T) {
	s, err := NewSourceFromFile("./source.go")
	if err != nil {
		t.Fatal(err)
	}

	if !strings.HasSuffix(s.Path, "source.go") {
		t.Errorf("Unexpected file name %s", s.Path)
	}

	if s.Code == nil {
		t.Errorf("Code was not read properly")
	}

	if !s.Exists {
		t.Errorf("File must exist")
	}
}

func TestUnexistFile(t *testing.T) {
	_, err := NewSourceFromFile("./__unknown_file.ml")
	if err == nil {
		t.Fatalf("Unknown error must cause an error")
	}
}

func TestReadFromStdin(t *testing.T) {
	s, err := NewSourceFromStdin()
	if err != nil {
		t.Fatal(err)
	}

	if s.Path != "<stdin>" {
		t.Errorf("Unexpected file name %s", s.Path)
	}

	if s.Code == nil {
		t.Errorf("Code was not read properly")
	}

	if s.Exists {
		t.Errorf("File must not exist")
	}
}

func TestBaseName(t *testing.T) {
	fromFile, err := NewSourceFromFile("./source.go")
	if err != nil {
		t.Fatal(err)
	}
	fromStdin, err := NewSourceFromStdin()
	if err != nil {
		t.Fatal(err)
	}
	fromDummy := NewDummySource("test")

	for _, tc := range []struct {
		expected string
		source   *Source
	}{
		{"source", fromFile},
		{"out", fromStdin},
		{"out", fromDummy},
	} {
		actual := tc.source.BaseName()
		if tc.expected != actual {
			t.Errorf("Expected base name of '%s' to be '%s', but actually it was '%s'", tc.source.Path, tc.expected, actual)
		}
	}
}

func TestSourceString(t *testing.T) {
	s := NewDummySource("")
	if s.String() != "source:<dummy>" {
		t.Fatal("Unknown source name:", s)
	}
}
