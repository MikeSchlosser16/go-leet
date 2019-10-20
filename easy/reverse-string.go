func reverseString(s []byte)  {
	for i := 0; i < len(s) / 2; i++{
		leadElement := s[i]
		endElement := s[len(s) - i - 1]
		s[i] = endElement
		s[len(s) - i - 1] = leadElement
	}
}