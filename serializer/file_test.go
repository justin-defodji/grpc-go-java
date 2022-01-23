package serializer

import (
	"github.com/stretchr/testify/require"
	"testing"

	"../sample"
)

func TestFIleSerializer(t *testing.T) {
	t.Parallel()

	binaryFIle := "../tmp/laptop.bin"

	laptop1 := sample.NewLaptop()

	err := WriteProtobufBinaryFile(laptop1, binaryFIle)
	require.NoError(t, err)
}
