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

func TestStringBuilder2(t *testing.T) {
	var builder StringBuilder

	builder.Appendf("%05d, %2.2f, %s", 1, 1.1, "テスト")
	if builder.String() != `00001, 1.10, テスト` {
		t.Errorf("Appendf is failed: %s vs 00001, 1.10, テスト", builder.String())
	} else {
		t.Log(builder.String())
	}
}
