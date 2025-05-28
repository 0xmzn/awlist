package cli

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetOutputWriter_EmptyPath(t *testing.T) {
	writer, closer, err := getOutputWriter("")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if writer != os.Stdout {
		t.Errorf("expected os.Stdout, got %v", writer)
	}
	if closer != nil {
		t.Errorf("expected nil closer for stdout, got non-nil: %T", closer)
	}

}

func TestGetOutputWriter_ValidFile(t *testing.T) {
	tmpDir := t.TempDir()
	outputPath := filepath.Join(tmpDir, "output.txt")

	writer, closer, err := getOutputWriter(outputPath)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if writer == nil {
		t.Fatal("expected writer, got nil")
	}
	if closer == nil {
		t.Fatal("expected closer, got nil")
	}
	defer closer()

	content := []byte("test")
	if _, err := writer.Write(content); err != nil {
		t.Fatalf("failed to write to file: %v", err)
	}

	data, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("failed to read written file: %v", err)
	}

	if string(data) != string(content) {
		t.Errorf("expected content %q, got %q", content, data)
	}
}

func TestGetOutputWriter_InvalidPath(t *testing.T) {
	_, _, err := getOutputWriter("/someinvalid/pathhy/output.xtx")
	if err == nil {
		t.Fatal("expected error for invalid path, got nil")
	}
}
