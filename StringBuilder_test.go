package utility

import (
	"testing"
)

func TestStringBuilder(t *testing.T) {
	var builder StringBuilder

	// Append
	builder.Append(`aaa`)
	if builder.String() != `aaa` {
		t.Error(`Append1 is failed`)
	}

	// Length
	if builder.Length() != len(`aaa`) {
		t.Error(`Length1 is failed`)
	}

	// Delete
	builder.Delete()
	if builder.String() != `` {
		t.Error(`Delete1 is failed`)
	}

	// Append
	builder.Append(`multiバイトも入っている状態`)
	if builder.String() != `multiバイトも入っている状態` {
		t.Error(`Append2 is failed`)
	}

	// Length
	if builder.Length() != len(`multiバイトも入っている状態`) {
		t.Error(`Length2 is failed`)
	}

	// Delete
	builder.Delete()
	if builder.String() != `` {
		t.Error(`Delete2 is failed`)
	}

	// Bytes
	builder.Append(`multiバイトも入っている状態`)
	if string(builder.Bytes()) != `multiバイトも入っている状態` {
		t.Error(`Bytes1 is failed`)
	}
}
