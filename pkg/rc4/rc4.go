package rc4

const STATE_ARRAY_LEN = 256

type State struct {
	i int
	j int
	S [STATE_ARRAY_LEN]byte
}

func Initialize() *State{
	return &State{
		i: 0,
		j: 0,
		S: [256]byte{},
	}
}

func Operate(plaintext, key []byte, state *State) []byte {
	S := (*state).S
	i := (*state).i
	j := (*state).j

	ciphertext := make([]byte, len(plaintext), len(plaintext))

	if S[0] == 0 && S[1] == 0 {
		for i := 0; i < 256; i++ {
			S[i] = byte(i)
		}

		for i := 0; i < 256; i++ {
			j = (j + int(S[i]) + int(key[i%len(key)])) % 256
			tmp := S[i]
			S[i] = S[j]
			S[j] = tmp
		}

		i = 0
		j = 0
	}

	ctIndex := 0
	ptIndex := 0

	for plaintextLen := len(plaintext); plaintextLen > 0; plaintextLen-- {
		i = (i+1)%256
		j = (j+int(S[i]))%256
		tmp := S[i]
		S[i] = S[j]
		S[j] = tmp
		ciphertext[ctIndex] = S[ ( int(S[i]) + int(S[j])  ) %256 ] ^ plaintext[ptIndex]
		ctIndex++
		ptIndex++
	}

	(*state).i = i
	(*state).j = j
	return ciphertext
}
