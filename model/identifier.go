package model

type BasicIdentifier string

func (b BasicIdentifier) String() string {
	return string(b)
}

func (b BasicIdentifier) IsValid() bool {
	for _, r := range b {
		if !allowedIDChars[r] {
			return false // Found a character not in the allowed list
		}
	}
	return true // All characters are in the allowed list
}

func (b BasicIdentifier) Ptr() *BasicIdentifier {
	return &b
}
