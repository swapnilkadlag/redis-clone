package main

import (
	"bufio"
	"errors"
	"strconv"
)

func readTillEOL(reader *bufio.Reader) (string, error) {
	s := ""
	var b byte
	var err error
	for {
		b, err = reader.ReadByte()
		if err != nil {
			return s, err
		}
		if b == cr {
			reader.ReadByte()
			break
		}
		s += string(b)
	}
	return s, nil
}

func readChar(reader *bufio.Reader) (rune, error) {
	b, err := reader.ReadByte()
	return rune(b), err
}

func readInt(reader *bufio.Reader) (int, error) {
	s, err := readTillEOL(reader)
	if err != nil {
		return -1, err
	}
	i, err := strconv.Atoi(s)
	return i, err
}

func readSimpleString(reader *bufio.Reader) (string, error) {
	return readTillEOL(reader)
}

func readBulkString(reader *bufio.Reader) (string, error) {
	_, err := readInt(reader)
	if err != nil {
		return "", err
	}
	s, err := readTillEOL(reader)
	return s, err
}

func readNext(reader *bufio.Reader) (any, error) {
	t, err := readChar(reader)
	if err != nil {
		return nil, err
	}
	switch t {
	case simpleString:
		return readSimpleString(reader)
	case bulkString:
		return readBulkString(reader)
	case integer:
		return readInt(reader)
	default:
		return nil, errors.New("Invalid data type")
	}
}

func readArray(reader *bufio.Reader) ([]any, error) {
	size, err := readInt(reader)
	if err != nil {
		return nil, err
	}
	if size == 0 {
		return nil, errors.New("Array size can't be 0")
	}
	arr := make([]any, size)
	for i := range arr {
		arr[i], err = readNext(reader)
		if err != nil {
			return arr, err
		}
	}
	return arr, nil
}
