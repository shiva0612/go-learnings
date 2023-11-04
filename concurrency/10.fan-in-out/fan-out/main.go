package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	ch1, err := read("file1.csv")
	if err != nil {
		panic(fmt.Errorf("Could not read file1 %v", err))
	}

	//-

	br1 := breakup("1", ch1)
	br2 := breakup("2", ch1)
	br3 := breakup("3", ch1)

	for {
		if br1 == nil && br2 == nil && br3 == nil {
			break
		}

		select {
		case _, open := <-br1:
			if !open {
				br1 = nil
			}
		case _, open := <-br2:
			if !open {
				br2 = nil
			}
		case _, open := <-br3:
			if !open {
				br3 = nil
			}
		}
	}

	fmt.Println("All completed, exiting")
}

func breakup(worker string, ch <-chan []string) chan struct{} {
	chE := make(chan struct{})

	go func() {
		defer close(chE)
		for v := range ch {
			fmt.Println(worker, v)
		}
	}()

	return chE
}

func read(file string) (<-chan []string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("opening file %v", err)
	}

	ch := make(chan []string)

	cr := csv.NewReader(f)

	go func() {
		for {
			record, err := cr.Read()
			if err == io.EOF {
				close(ch)

				return
			}

			ch <- record
		}
	}()

	return ch, nil
}
