package output

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CSVWriter struct {
}

type CSVData interface {
	CSVHeader() []string
	CSVFormat() []string
}

func (w CSVWriter) ToFile(filepath string, data []CSVData) error {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return fmt.Errorf("error opening output file: %w", err)
	}
	writer := csv.NewWriter(file)

	// Write Headers
	err = writer.Write(data[0].CSVHeader())
	if err != nil {
		return fmt.Errorf("error writing header %s to csv: %w", data[0].CSVHeader(), err)
	}

	// Write Rows
	for _, entry := range data {
		err = writer.Write(entry.CSVFormat())
		if err != nil {
			return fmt.Errorf("error writing record %s to csv: %w", entry.CSVFormat(), err)
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		_ = file.Close()
		return fmt.Errorf("error writing CSV: %w", err)
	}
	if err := file.Close(); err != nil {
		return fmt.Errorf("error closing file: %w", err)
	}
	return nil
}
