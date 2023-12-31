package cmd

import (
	"camm_extractor/internal/decoder"
	"fmt"
	"github.com/spf13/cobra"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decodes a binary CAMM data file",
	Long:  `Expects a single argument, the input binary file`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return callback(args[0])
	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)
}

func callback(binaryFile string) error {
	cammStream, err := decoder.Run(binaryFile)
	if err != nil {
		return err
	}
	for _, packet := range cammStream.DecodedPackets {
		fmt.Printf("%d | %s\n", packet.PacketType(), packet)
	}
	return nil
}
