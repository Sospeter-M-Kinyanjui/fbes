package main

import (
	"flag"
	"fmt"
	"fuck/fconcat"
	"fuck/fcopy"
	"fuck/fdelete"
	"fuck/fmore"
	"fuck/fmove"
	"fuck/frename"
)

func init() {
	fmt.Println("* - _ Welcome to f**k File Manager _ - *")
}

func main() {
	f_copy := flag.String("copy", "copy", "flag for copying files")
	f_move := flag.String("move", "move", "flag for moving files")
	f_more := flag.String("more", "more", "flag for displaying file contents")
	f_delete := flag.String("delete", "delete", "flag to delete file")
	f_rename := flag.String("rename", "rename", "flag to rename file")
	f_concat := flag.String("concat", "concat", "flag to concatenate two files contents")

	flag.Parse()

	if *f_copy == "copy" {
		fcopy.Fcopy()
	}

	if *f_move == "move" {
		fmove.Fmove()
	}

	if *f_more == "more" {
		fmore.Fmore()
	}

	if *f_rename == "rename" {
		frename.Frename()
	}

	if *f_concat == "concat" {
		fconcat.Fconcat()
	}

	if *f_delete == "delete" {
		fdelete.Fdelete()
	}

}
