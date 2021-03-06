package euclid;
import org.junit.*;
import junit.framework.TestCase;

public class TestEuclid extends TestCase {

	/**
	 * This class houses the input/expected output pairs for the function
	 * Euclid().
	 */
	private class TestCase {
		public int input1;
		public int input2;
		public int output;

		public TestCase(int input1, int input2, int output) {
			this.input1=input1;
			this.input2=input2;
			this.output=output;
		}
	}

	/**
	 * A table of test cases.
	 */
	TestCase[] testCases = {
		new TestCase(20,100,20),
		new TestCase(100,20,20),
		new TestCase(8,15,1),
		new TestCase(15,8,1),
		new TestCase(21,133,7),
		new TestCase(133,21,7),
		new TestCase(36,4352,4),
		new TestCase(4352,36,4),
		new TestCase(-4352,36,-1),
		new TestCase(4352,-36,-1),
		new TestCase(-21,-133,-1),
		new TestCase(0,133,0),
		new TestCase(21,0,0),
		new TestCase(0,0,0)
	};

	/**
	 * A table driven unit test for the function Euclid.
	 */
	@Test
	public void testOutput() {
		for(int i = 0; i < testCases.length; i++) {
			int result = Euclid.GCD(testCases[i].input1,testCases[i].input2);
			assertTrue("Case "+i+ " Failed with Output: "+result,result==testCases[i].output);
		}
	}
}
