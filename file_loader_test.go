package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func Test_loadMapJSONFile(t *testing.T) {
	fileWithValidJSONContent := "/tmp/enviin-test-property-valid.json"
	fileWithInvalidJSONContent := "/tmp/enviin-test-property-invalid.json"

	defer func() {
		for _, filepath := range []string{fileWithInvalidJSONContent, fileWithValidJSONContent} {
			if _, err := os.Stat(filepath); err == nil {
				os.Remove(filepath)
			}
		}
	}()

	validJsonContent := []byte(`
		{
			"simpleString" :  "this is string"
		}
	`)
	ioutil.WriteFile(fileWithValidJSONContent, validJsonContent, 0644)
	ioutil.WriteFile(fileWithInvalidJSONContent, []byte("hello"), 0664)

	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{"valid Json", args{filepath: fileWithValidJSONContent}, map[string]interface{}{"simpleString": toI("this is string")}, false},
		{"invalid Json", args{filepath: fileWithInvalidJSONContent}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadMapJSONFile(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadMapJSONFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadMapJSONFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadMapYAMLFile(t *testing.T) {
	fileWithValidYAMLContent := "/tmp/enviin-test-property-valid.yaml"
	fileWithInvalidYAMLContent := "/tmp/enviin-test-property-invalid.yaml"

	defer func() {
		for _, filepath := range []string{fileWithValidYAMLContent, fileWithInvalidYAMLContent} {
			if _, err := os.Stat(filepath); err == nil {
				os.Remove(filepath)
			}
		}
	}()

	validYAMLContent := []byte(`attribute: Easy!`)
	ioutil.WriteFile(fileWithValidYAMLContent, validYAMLContent, 0644)
	ioutil.WriteFile(fileWithInvalidYAMLContent, []byte("hello"), 0664)

	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{"valid YAML", args{filepath: fileWithValidYAMLContent}, map[string]interface{}{"attribute": toI("Easy!")}, false},
		{"invalid YAML", args{filepath: fileWithInvalidYAMLContent}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadMapYAMLFile(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadMapYAMLFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadMapYAMLFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
