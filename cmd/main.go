package main

import (
	"bufio"
	"log"
	"os"

	"diego.pizza/auction-house/pkg/bid"
	"diego.pizza/auction-house/pkg/engine"
	"diego.pizza/auction-house/pkg/hearthbeat"
	"diego.pizza/auction-house/pkg/listing"
	"diego.pizza/auction-house/pkg/runner"
)

func main() {
	if err := run(os.Stdin, os.Stdout, os.Stderr); err != nil {
		log.Fatal(err)
	}
}

func run(stdin *os.File, stdout *os.File, stderr *os.File) error {
	scanner := bufio.NewScanner(stdin)
	runner := runner.New(bid.New(), listing.New(), hearthbeat.New(), engine.New())

	for scanner.Scan() {
		iter, err := runner.Run(scanner.Text())
		if err != nil {
			stderr.WriteString(err.Error())
			stdout.WriteString("\n")
		}
		for line, next := iter(); next; line, next = iter() {
			stdout.WriteString(line)
			stdout.WriteString("\n")
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
