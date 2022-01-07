package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	testCases := []struct {
		desc  string
		key   string
		query []byte
		err   error
	}{
		{
			desc:  "a",
			key:   "micro",
			query: []byte(`{"service":{}}`),
			err:   nil,
		},
		{
			desc:  "b",
			key:   "micro",
			query: []byte(``),
			err:   fmt.Errorf("unexpected end of JSON input"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := Confs.Set(tc.key, tc.query)
			if err != nil {
				assert.Equal(t, err.Error(), tc.err.Error())
			}
		})
	}
}

func TestGet(t *testing.T) {
	testCases := []struct {
		desc  string
		debug bool
	}{
		{
			desc:  "a",
			debug: false,
		},
		{
			desc:  "b",
			debug: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			Confs.SetDebug(tC.debug)
			cnfs := Confs
			if cnfs.Debug != tC.debug {
				assert.Equal(t, tC.debug, cnfs.Debug)
			}
		})
	}
}

func TestGetDebug(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "a",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			Confs.GetDebug()
		})
	}
}

func TestLoad(t *testing.T) {
	testCases := []struct {
		desc, path string
		err        error
		data       string
		f          func(name, data string)
	}{
		{
			desc: "a",
			path: "config.yaml",
			err:  nil,
			data: ``,
			f: func(name, data string) {
				ioutil.WriteFile(name, []byte(data), 0755)
				go func() {
					time.Sleep(time.Millisecond * 100)
					os.Remove(name)
				}()
			},
		},
		{
			desc: "b",
			path: "config.yaml",
			err:  nil,
			data: `{"service":{}}`,
			f: func(name, data string) {
				ioutil.WriteFile(name, []byte(data), 0755)
				go func() {
					time.Sleep(time.Millisecond * 100)
					os.Remove(name)
				}()
			},
		},
		{
			desc: "c",
			path: "",
			err:  fmt.Errorf("file not exists"),
			data: `{"service":{}}`,
			f:    func(name, data string) {},
		},
		{
			desc: "d",
			path: "",
			err:  fmt.Errorf("file not exists"),
			data: ``,
			f: func(name, data string) {
				// ioutil.WriteFile(name, []byte(data), 0755)
				go func() {
					time.Sleep(time.Millisecond * 100)
					os.Remove(name)
				}()
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			tc.f(tc.path, tc.data)
			if err := Confs.Load(tc.path); err != nil {
				assert.Equal(t, tc.err.Error(), err.Error())
			}
			time.Sleep(time.Millisecond * 150)
		})
	}
}

func TestFile(t *testing.T) {
	testCases := []struct {
		desc, path string
		err        error
		data       string
		f          func(name, data string)
	}{
		{
			desc: "a",
			path: "config.yaml",
			err:  nil,
			data: ``,
			f: func(name, data string) {
				ioutil.WriteFile(name, []byte(data), 0755)
				go func() {
					time.Sleep(time.Millisecond * 100)
					os.Remove(name)
				}()
			},
		},
		{
			desc: "b",
			path: "config.yaml",
			err:  nil,
			data: `{"service":{}}`,
			f: func(name, data string) {
				ioutil.WriteFile(name, []byte(data), 0755)
				go func() {
					time.Sleep(time.Millisecond * 100)
					os.Remove(name)
				}()
			},
		},
		{
			desc: "c",
			path: "",
			err:  fmt.Errorf("open config.yaml: no such file or directory"),
			data: `{"service":{}}`,
			f:    func(name, data string) {},
		},
		{
			desc: "d",
			path: "path.txt",
			err:  fmt.Errorf("Unsupported Config Type \"txt\""),
			data: ``,
			f: func(name, data string) {
				ioutil.WriteFile(name, []byte(data), 0755)
				go func() {
					time.Sleep(time.Millisecond * 100)
					os.Remove(name)
				}()
			},
		},
		{
			desc: "e",
			path: "config.yaml",
			err:  fmt.Errorf("While parsing config: yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `\\n\\t{\\n...` into map[string]interface {}"),
			data: `\n\t{\n"service":\n{}\n}\n`,
			f: func(name, data string) {
				ioutil.WriteFile(name, []byte(data), 0755)
				go func() {
					time.Sleep(time.Millisecond * 100)
					os.Remove(name)
				}()
			},
		},
		{
			desc: "f",
			path: "config.yaml",
			err:  fmt.Errorf("x"),
			data: ``,
			f: func(name, data string) {
				ioutil.WriteFile(name, []byte(data), 0755)
				go func() {
					time.Sleep(time.Millisecond * 100)
					os.Remove(name)
				}()
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			tst := Config{}
			tc.f(tc.path, tc.data)
			if err := tst.file(tc.path); err != nil {
				assert.Equal(t, tc.err.Error(), err.Error())
			}
			time.Sleep(time.Millisecond * 150)
		})
	}
}
