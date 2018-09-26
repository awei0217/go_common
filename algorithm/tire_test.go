package algorithm

import "testing"

func TestTireTree_Insert(t *testing.T) {
	tt := &TireTree{}
	tt.Insert("abbcd")
	tt.Insert("cd")
	tt.Insert("abr")
	tt.Insert("abcd")
	t.Log(tt.Query("cd"))
	t.Log(tt.Query("wer"))
	//tt.PreTraverse(tt.root)
	t.Log(tt.CountPrefix("c"))
}