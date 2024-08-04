package main

import "google.golang.org/protobuf/types/known/timestamppb"

func main() {

	timestamppb.Now().AsTime()

}
