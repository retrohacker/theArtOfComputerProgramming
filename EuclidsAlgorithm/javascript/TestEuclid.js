//A value of input/output pairs for a unit test
var testCases = [
	[20,100,20],
	[100,20,0],
	[8,15,1],
	[15,8,0],
	[21,133,7],
	[133,21,0],
	[36,4352,4],
	[4352,36,0],
	[-8,15,-1],
	[8,-15,-1],
	[-8,-15,-1],
	[0,36,0],
	[4352,0,0],
	[0,0,0]
];

/**
 * This function will act as a unit test of Euclid's algorithm. It uses a table
 * driven system for testing values.
 */
function TestEuclid() {
	passed = true;
	for(i=0;i<testCases.length;i++) {
		if(!verify(testCases[i],i)) {
			passed = false;
		}
	}	
	if(passed) {
		document.write("OK (Tests Passed)");
	} else {
		document.write("FAIL (One or more tests failed)");
	}
}

/**
 * This function verify the output of Euclid's algorithm with its testcase
 * data. It accepts two values Number[3] and Number. Number[3] should be in
 * the form [input1,input2,output]. Number is the test number being run and is
 * used soely for error reporting.
 */
function verify(testCase, caseNum) {
	result = Euclid(testCase[0],testCase[1])
	if(result!=testCase[2]) {
		document.write("Case "+caseNum+" FAILED with Output: "+result+" Expected: "+testCase[2]+"<br/>");
		return false;
	}
	return true;
}
