testCases = [
		[1769,551,5,-16,29],
		[-1769,551,0,0,-1],
		[1769,-551,0,0,-1],
		[-1769,-551,0,0,-1],
		[0,551,0,0,0],
		[0,0,0,0,0],
		[0,551,0,0,0],
		[551,1769,-16,5,29],
		[120,23,-9,47,1],
		[23,120,47,-9,1],
		[46,57,-26,21,1],

]

function TestEuclidExtended() {
	passed=true
	for(i=0;i<testCases.length;i++) {
		result=EuclidExtended(testCases[i][0],testCases[i][1]);
		if(result[0]!=testCases[i][2]||result[1]!=testCases[i][3]||result[2]!=testCases[i][4]) {
			passed=false
			document.write("Test "+i+" FAILED Expected: "+testCases[i].slice(2)+" Returned: "+result+"<br />");
		}
	}
	if(passed) {
		document.write("OK (Test Passed)");
	} else {
		document.write("Test Failed");
	}
}
