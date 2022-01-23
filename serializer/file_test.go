package serializer_test

import (
	"github.com/stretchr/testify/require"
	"testing"

	"gitlab.com/techschool/pcbook/sample"

	"../serializer"
)

func TestFIleSerializer(t *testing.T) {
	t.Parallel()

	binaryFIle := "../tmp/laptop.bin"

	laptop1 := sample.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFIle)
	require.NoError(t, err)
}
