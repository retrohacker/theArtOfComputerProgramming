var testCases = [
	[20,100,20],
	[100,20,20],
	[8,15,1],
	[15,8,1],
	[21,133,7],
	[133,21,7],
	[36,4352,4],
	[4352,36,4]
];

function TestEuclid() {
	passed = true;
	for(i=0;i<testCases.length;i++) {
		if(!verify(testCases[i],i)) {
			passed = false;
			break;
		}
	}	
	if(passed) {
		document.write("OK (Tests Passed)");
	} else {
		document.write("Tests FAILED");
	}
}

function verify(testCase, caseNum) {
	result = Euclid(testCase[0],testCase[1])
	if(result!=testCase[2]) {
		document.write("Case "+caseNum+" FAILED with Output: "+result+"<br/>");
		return false;
	}
	return true;
}
