package output

type FileWriter interface {
	ToFile(options WriterOptions)
}

type WriterOptions struct {
	OutputFilepath string
	EnabledPackets []int
}

//// Define which types the user wants to keep
//userSelectedTypes := map[int]bool{
//1: true,  // Assuming type 1 is of interest
//3: true,  // Assuming type 3 is of interest
//// Add more types as required
//}
//
//// Call the filtering and writing function
//err := cammStream.FilterAndWriteToFile(userSelectedTypes, "output.csv")
//if err != nil {
//log.Fatalf("Failed to write filtered data to file: %v", err)
//}
//}
