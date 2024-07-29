package variables

func DeclareString(s string) string {
	var output string = s
	return output
}

func ShortDeclareString(s string) string {
	output := s
	return output
}

func DeclareThreeStrings(s1, s2, s3 string) (string, string, string) {
	var (
		output1 string
		output2 string
		output3 string
	)
	output1, output2, output3 = s1, s2, s3
	return output1, output2, output3
}

func DeclareThreeStringsOneLine(s1, s2, s3 string) (string, string, string) {
	var output1, output2, output3 string = s1, s2, s3
	return output1, output2, output3
}

func DeclareThreeStringsOneLineMixedType(a1, a2, a3 any) (any, any, any) {
	var output1, output2, output3 = a1, a2, a3
	return output1, output2, output3
}

func ShortDeclareThreeStrings(s1, s2, s3 string) (string, string, string) {
	output1, output2, output3 := s1, s2, s3
	return output1, output2, output3
}
